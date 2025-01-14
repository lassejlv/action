#!/bin/bash

# Update script for actionfile

# Check if actionfile exists
if [ ! -f "/usr/local/bin/action" ]; then
    printf "❌ actionfile is not installed. Please run install.sh first\n"
    exit 1
fi

# Get current version
currentVersion=$(/usr/local/bin/action --version)
printf "Current version: $currentVersion\n"

# Get latest version from GitHub
latestUrl=$(curl -s https://api.github.com/repos/lassejlv/action/releases/latest | grep "browser_download_url" | cut -d '"' -f 4)
latestVersion=$(echo $latestUrl | cut -d '/' -f 8)
printf "Latest version: $latestVersion\n"

# Compare versions
if [ "$currentVersion" == "$latestVersion" ]; then
    printf "✅ You are already running the latest version!\n"
    exit 0
fi

# Determine system type
userSystem=$(uname -s)
userArch=$(uname -m)
printf "Detected system: $userSystem $userArch\n"

# Set the correct download URL based on system and architecture
if [ "$userSystem" == "Linux" ]; then
    if [ "$userArch" == "x86_64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$latestVersion/actionfile-linux-x64"
    elif [ "$userArch" == "aarch64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$latestVersion/actionfile-linux-arm64"
    fi
elif [ "$userSystem" == "Darwin" ]; then
    if [ "$userArch" == "x86_64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$latestVersion/actionfile-mac-x64"
    elif [ "$userArch" == "arm64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$latestVersion/actionfile-mac-arm64"
    fi
elif [ "$userSystem" == "Windows_NT" ]; then
    downloadUrl="https://github.com/lassejlv/action/releases/download/$latestVersion/actionfile-win-x64.exe"
fi

if [ -z "$downloadUrl" ]; then
    printf "Error: Unsupported system: $userSystem $userArch\n"
    exit 1
fi

printf "Downloading update from: $downloadUrl\n"

# Create temporary directory for download
tmpdir=$(mktemp -d)
curl -L -o "$tmpdir/action" "$downloadUrl"

# Check if download was successful
if [ $? -ne 0 ]; then
    printf "❌ Download failed\n"
    rm -rf "$tmpdir"
    exit 1
fi

# Update the binary with sudo
printf "Installing update to /usr/local/bin (requires sudo)...\n"
sudo mv "$tmpdir/action" /usr/local/bin/
sudo chmod +x /usr/local/bin/action

# Check if update was successful
if [ $? -eq 0 ]; then
    printf "✅ Update successful!\n"
    printf "Updated from $currentVersion to $latestVersion\n"
else
    printf "❌ Update failed\n"
    exit 1
fi

# Cleanup
rm -rf "$tmpdir"

exit 0
