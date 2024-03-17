#!/bin/bash

sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install git -y

# Clone the repository
mkdir -p /tmp/weeserve
cd /tmp/weeserve
URL="https://git.serverify.de/SecNex/weeserve.git"
git clone $URL

# Create the directory for the binary file
mkdir -p /usr/local/bin/weeserve

# Check if architecture is 64-Bit and linux
if [ $(uname -m) == "x86_64" ]; then
    # Copy the binary file to the directory
    cp /tmp/weeserve/weeserve/build/weeserve_linux_amd64 /usr/local/bin/weeserve/weeserve
    exit 0
fi

# Check if architecture is 32-Bit and linux
if [ $(uname -m) == "i686" ]; then
    # Copy the binary file to the directory
    cp /tmp/weeserve/weeserve/build/weeserve_linux_386 /usr/local/bin/weeserve/weeserve
    exit 0
fi

# Check if architecture is ARM64 and linux
if [ $(uname -m) == "aarch64" ]; then
    # Copy the binary file to the directory
    cp /tmp/weeserve/weeserve/build/weeserve_linux_arm64 /usr/local/bin/weeserve/weeserve
    exit 0
fi

# Check if architecture is AMD64 and macOS
if [ $(uname -m) == "x86_64" ] && [ $(uname) == "Darwin" ]; then
    # Copy the binary file to the directory
    cp /tmp/weeserve/weeserve/build/weeserve_darwin_amd64 /usr/local/bin/weeserve/weeserve
    exit 0
fi

# Check if architecture is ARM64 and macOS
if [ $(uname -m) == "arm64" ] && [ $(uname) == "Darwin" ]; then
    # Copy the binary file to the directory
    cp /tmp/weeserve/weeserve/build/weeserve_darwin_arm64 /usr/local/bin/weeserve/weeserve
    exit 0
fi