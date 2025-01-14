#!/bin/bash

# Uninstall script for actionfile

# Check if actionfile exists
if [ ! -f "/usr/local/bin/action" ]; then
    printf "❌ actionfile is not installed in /usr/local/bin\n"
    exit 1
fi

# Remove the binary with sudo
printf "Removing action from /usr/local/bin (requires sudo)...\n"
sudo rm /usr/local/bin/action

# Check if removal was successful
if [ $? -eq 0 ]; then
    printf "✅ Uninstallation successful!\n"
    printf "actionfile has been removed from your system.\n"
else
    printf "❌ Uninstallation failed\n"
    exit 1
fi

exit 0
