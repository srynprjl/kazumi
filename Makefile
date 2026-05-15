# Project variables
BINARY_NAME=kazumi
VERSION=alpha
BUILD_DIR=bin
INSTALL_DIR=$(HOME)/.local/bin

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags="-s -w -extldflags '-static'"

.PHONY: all build clean test fmt install uninstall help

all: fmt build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build:
	@echo "Building $(BINARY_NAME) $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=1 $(GOBUILD) $(LDFLAGS) -trimpath -o $(BUILD_DIR)/$(BINARY_NAME) cmd/kazumi/main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@$(GOCLEAN)
	@echo "Cleaned."

test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

fmt:
	@echo "Formatting code..."
	@$(GOFMT) ./...

tidy:
	@echo "Tidying modules..."
	@$(GOMOD) tidy

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@mkdir -p $(INSTALL_DIR)
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installation complete."

uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)..."
	@rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstalled."
