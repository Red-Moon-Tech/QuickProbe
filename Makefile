# Makefile for Go Application

# Variables
APP_NAME := qprobe     # Change to your application name
SRC := ./src/    # Change this to your application's source files
BUILD_DIR := ./build
GOOS ?= linux         # Default target OS
GOARCH ?= amd64       # Default target architecture

# Ensure the build directory exists
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Target for building the application
.PHONY: build
build: $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC)

# Clean up build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Target for building for Windows (amd64)
.PHONY: build-windows
build-windows: GOOS := windows
build-windows: GOARCH := amd64
build-windows: build

# Target for building for Linux (amd64)
.PHONY: build-linux
build-linux: GOOS := linux
build-linux: GOARCH := amd64
build-linux: build

# Target for building for macOS (amd64)
.PHONY: build-macos
build-macos: GOOS := darwin
build-macos: GOARCH := amd64
build-macos: build

# Define a target to build for a custom OS and architecture
.PHONY: build-custom
build-custom:
	@echo "Building for OS: $(GOOS), ARCH: $(GOARCH)"
	$(MAKE) build

# Default target
.PHONY: all
all: build

