#!/bin/bash

# gourl Installer Script
# Installs gourl directly from GitHub

set -e

REPO="ram-ai-kumar/gourl"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
BINARY_NAME="${BINARY_NAME:-gourl}"

# Parse arguments
STABLE=false
BRANCH="develop"
INSTALL_TYPE="edge"

while [[ "$#" -gt 0 ]]; do
    case $1 in
        --stable) STABLE=true; BRANCH="main"; INSTALL_TYPE="stable" ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

echo "🚀 Installing gourl ($INSTALL_TYPE) from branch '$BRANCH'..."

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

if [ "$STABLE" = true ]; then
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
    echo "📦 Fetching latest stable release..."
    LATEST_RELEASE=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' || echo "")

    if [ -z "$LATEST_RELEASE" ]; then
        echo "⚠️  Could not fetch latest release. Proceeding with source build (fallback)..."
        STABLE=false
    else
        echo "📥 Downloading gourl $LATEST_RELEASE ($OS-$ARCH)..."
        DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/gourl-$OS-$ARCH"
        if curl -fsSL "$DOWNLOAD_URL" -o "$INSTALL_DIR/$BINARY_NAME"; then
            chmod +x "$INSTALL_DIR/$BINARY_NAME"
        else
            echo "❌ Binary download failed. Proceeding with source build (fallback)..."
            STABLE=false
        fi
    fi
fi

if [ "$STABLE" = false ]; then
    # Check for dependencies
    if ! command -v git &> /dev/null; then
        echo "❌ Error: 'git' is not installed. Required for source build."
        exit 1
    fi
    if ! command -v go &> /dev/null; then
        echo "❌ Error: 'go' is not installed. Required for source build."
        exit 1
    fi

    # Install from source
    TEMP_DIR=$(mktemp -d)
    echo "📥 Cloning branch '$BRANCH'..."
    git clone --branch "$BRANCH" --single-branch "https://github.com/$REPO.git" "$TEMP_DIR"

    echo "🔨 Building gourl..."
    cd "$TEMP_DIR"
    # Set version during build
    VERSION_TAG="edge-$(date +%Y%m%d)"
    if [ "$BRANCH" = "main" ]; then VERSION_TAG="stable-$(date +%Y%m%d)"; fi
    
    go build -ldflags "-X main.Version=$VERSION_TAG" -o "$INSTALL_DIR/$BINARY_NAME"
    
    cd /
    rm -rf "$TEMP_DIR"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
fi

# Check if install directory is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "⚠️  $INSTALL_DIR is not in your PATH"
    echo "Add the following line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo "export PATH=\"\$PATH:$INSTALL_DIR\""
    echo "Then restart your shell or run: source ~/.bashrc"
fi

echo "✅ gourl installed successfully!"
echo "Run 'gourl version' to check the installed version."
