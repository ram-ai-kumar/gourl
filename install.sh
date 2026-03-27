#!/bin/bash

# gourl Installer Script
# Installs gourl directly from GitHub

set -e

REPO="ram-ai-kumar/gourl"
INSTALL_DIR="$HOME/.local/bin"
BINARY_NAME="gourl"

echo "🚀 Installing gourl from $REPO..."

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | tr '[:upper:]' '[:lower:]')

# Map architecture names
case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    armv7l) ARCH="arm" ;;
    *) echo "❌ Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Map OS names
case $OS in
    darwin) OS="darwin" ;;
    linux) OS="linux" ;;
    *) echo "❌ Unsupported OS: $OS"; exit 1 ;;
esac

# Get latest release version
echo "📦 Fetching latest release..."
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1')

if [ -z "$LATEST_RELEASE" ]; then
    echo "❌ Could not fetch latest release. Please check your internet connection."
    exit 1
fi

echo "📥 Downloading gourl $LATEST_RELEASE for $OS-$ARCH..."

# Download the binary
DOWNLOAD_URL="https://github.com/ram-ai-kumar/gourl/releases/download/$LATEST_RELEASE/gourl-$OS-$ARCH"
curl -L "$DOWNLOAD_URL" -o "$INSTALL_DIR/$BINARY_NAME"

# Make binary executable
chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Check if install directory is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "⚠️  $INSTALL_DIR is not in your PATH"
    echo "Add the following line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo "export PATH=\"\$PATH:$INSTALL_DIR\""
    echo "Then restart your shell or run: source ~/.bashrc"
fi

echo "✅ gourl installed successfully!"
echo "Run 'gourl help' to get started."
