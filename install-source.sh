#!/bin/bash

# gourl Installer Script - Build from Source
# Installs gourl by building from source code

set -e

REPO="ram-ai-kumar/gourl"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
BINARY_NAME="${BINARY_NAME:-gourl}"
TEMP_DIR=$(mktemp -d)

echo "🚀 Installing gourl from $REPO..."

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Clone and build
BRANCH=${1:-develop}
echo "📥 Cloning repository (branch: $BRANCH)..."
git clone --branch "$BRANCH" --single-branch https://github.com/ram-ai-kumar/gourl.git "$TEMP_DIR"

echo "🔨 Building gourl..."
cd "$TEMP_DIR"
go build -o "$INSTALL_DIR/$BINARY_NAME"

# Cleanup
cd /
rm -rf "$TEMP_DIR"

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
