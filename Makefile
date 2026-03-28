# Makefile for gourl

.PHONY: build test test-features clean install

# Build the gourl binary
build:
	go build -o gourl .

# Run unit tests
test:
	go test ./...

# Run feature tests with godog
test-features:
	./test-features.sh

# Install godog if not present and run feature tests
test-features-deps:
	@if ! command -v godog &> /dev/null; then \
		echo "Installing godog..."; \
		go install github.com/cucumber/godog/cmd/godog@latest; \
	fi
	godog

# Clean build artifacts
clean:
	rm -f gourl
	rm -f gourl-*
	go clean

# Install the binary
install: build
	cp gourl $(GOPATH)/bin/ || cp gourl ~/go/bin/ || echo "Please add gourl to your PATH manually"

# Format code
fmt:
	go fmt ./...

# Run all tests
test-all: test test-features

# Development setup
dev-setup:
	go install github.com/cucumber/godog/cmd/godog@latest
	go mod tidy
	go mod download
