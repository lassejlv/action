#!/bin/bash

# A script that installs the latest release for the user's system
# from github.com/lassejlv/action/releases/latest

# 1. set a variable for the latest release
latestUrl=$(curl -s https://api.github.com/repos/lassejlv/action/releases/latest | grep "browser_download_url" | cut -d '"' -f 4)
version=$(echo $latestUrl | cut -d '/' -f 8)
printf "Found version: $version\n"

# 2. determine the user's system
userSystem=$(uname -s)
userArch=$(uname -m)
printf "Detected system: $userSystem $userArch\n"

# Set the correct download URL based on system and architecture
if [ "$userSystem" == "Linux" ]; then
    if [ "$userArch" == "x86_64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$version/actionfile-linux-x64"
    elif [ "$userArch" == "aarch64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$version/actionfile-linux-arm64"
    fi
elif [ "$userSystem" == "Darwin" ]; then
    if [ "$userArch" == "x86_64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$version/actionfile-mac-x64"
    elif [ "$userArch" == "arm64" ]; then
        downloadUrl="https://github.com/lassejlv/action/releases/download/$version/actionfile-mac-arm64"
    fi
elif [ "$userSystem" == "Windows_NT" ]; then
    downloadUrl="https://github.com/lassejlv/action/releases/download/$version/actionfile-win-x64.exe"
fi

if [ -z "$downloadUrl" ]; then
    printf "Error: Unsupported system: $userSystem $userArch\n"
    exit 1
fi

printf "Downloading from: $downloadUrl\n"

# 3. download the file and move it to the user's bin folder
# Create a temporary directory for downloading
tmpdir=$(mktemp -d)
curl -L -o "$tmpdir/action" "$downloadUrl"

# Check if curl was successful
if [ $? -ne 0 ]; then
    printf "Error: Download failed\n"
    rm -rf "$tmpdir"
    exit 1
fi

# Move to /usr/local/bin with sudo
printf "Installing to /usr/local/bin (requires sudo)...\n"
sudo mv "$tmpdir/action" /usr/local/bin/
sudo chmod +x /usr/local/bin/action

# 4. check if the installation was successful
if [ $? -eq 0 ]; then
    printf "✅ Installation successful!\n"
    printf "You can now run: action --help\n"
else
    printf "❌ Installation failed\n"
fi

# 5. cleanup
rm -rf "$tmpdir"

exit 0
