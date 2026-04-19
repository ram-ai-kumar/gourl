# Makefile for gourl

.PHONY: build test test-features clean install fmt dev-setup

# Build the gourl binary
build:
	mkdir -p bin
	GOMODCACHE=$(shell pwd)/.cache/gomodcache go build -o bin/gourl ./cmd/gourl

# Run unit tests
test:
	GOMODCACHE=$(shell pwd)/.cache/gomodcache go test ./internal/...

# Run feature tests with godog
test-features: build
	./scripts/test-features.sh

# Install the binary
install: build
	cp bin/gourl $(GOPATH)/bin/ || cp bin/gourl ~/go/bin/ || echo "Please add gourl to your PATH manually"

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f gourl-*
	GOMODCACHE=$(shell pwd)/.cache/gomodcache go clean

# Format code
fmt:
	GOMODCACHE=$(shell pwd)/.cache/gomodcache go fmt ./...

# Run all tests
test-all: test test-features

# Development setup
dev-setup:
	GOMODCACHE=$(shell pwd)/.cache/gomodcache go install github.com/cucumber/godog/cmd/godog@latest
	GOMODCACHE=$(shell pwd)/.cache/gomodcache go mod tidy
	GOMODCACHE=$(shell pwd)/.cache/gomodcache go mod download
