# Makefile for building a Go binary

# Variables
APP_NAME := mawinter-gemini-advisor
SRC_DIR := .
BUILD_DIR := bin
GO := go

# Build the Go binary (static build)
.PHONY: bin
bin:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 $(GO) build -tags netgo -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

.PHONY: build
build:
	docker build -t azuki774/mawinter-gemini-advisor -f build/Dockerfile .

# Clean up build artifacts
.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

# Run the application
.PHONY: run
run: build
	./$(BUILD_DIR)/$(APP_NAME)

# Format the code
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Run tests
.PHONY: test
test:
	$(GO) test -v ./...

# Lint the code (using staticcheck)
.PHONY: lint
lint:
	staticcheck ./...
