.PHONY: build clean install test

BINARY_DIR = plugins/modules
SRC_DIR = src

# Build all modules
build:
	@echo "Building Go modules..."
	@mkdir -p $(BINARY_DIR)
	cd $(SRC_DIR) && go build -o ../$(BINARY_DIR)/hello_world ./hello_world/
	cd $(SRC_DIR) && go build -o ../$(BINARY_DIR)/file_manager ./file_manager/
	@echo "Build complete!"

# Clean built binaries
clean:
	@echo "Cleaning binaries..."
	rm -f $(BINARY_DIR)/*
	@echo "Clean complete!"

# Install collection locally
install: build
	ansible-galaxy collection install . --force

# Test modules
test: install
	ansible-playbook playbooks/test_modules.yml

# Development dependencies
deps:
	cd $(SRC_DIR) && go mod tidy

# Cross-compile for different platforms
build-all:
	@echo "Cross-compiling for multiple platforms..."
	@mkdir -p $(BINARY_DIR)

	# Linux amd64
	cd $(SRC_DIR) && GOOS=linux GOARCH=amd64 go build -o ../$(BINARY_DIR)/hello_world ./hello_world/
	cd $(SRC_DIR) && GOOS=linux GOARCH=amd64 go build -o ../$(BINARY_DIR)/file_manager ./file_manager/

	@echo "Cross-compilation complete!"