#!/bin/bash

# Test script for running godog feature tests
# This script builds the gourl binary and runs all cucumber feature tests

set -e

echo "🧪 Running gourl feature tests with godog..."

# Check if godog is available
if ! command -v godog &> /dev/null; then
    echo "📦 Installing godog..."
    go install github.com/cucumber/godog/cmd/godog@latest
fi

# Build the main binary first
echo "🔨 Building gourl binary..."
go build -o gourl .

# Run the feature tests
echo "🚀 Running feature tests..."
godog

echo "✅ All feature tests completed!"
