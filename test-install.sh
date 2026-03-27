#!/bin/bash

# Test version of install script that builds locally
set -e

INSTALL_DIR="$HOME/.local/bin"
BINARY_NAME="gourl"

echo "🚀 Building gourl locally..."

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Build the binary
go build -o "$INSTALL_DIR/$BINARY_NAME"

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
