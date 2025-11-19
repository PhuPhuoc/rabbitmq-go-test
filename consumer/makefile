.DEFAULT_GOAL := help

# App info
APP_NAME := rabbitmq-go-test-consumer
BIN_DIR := bin
MAIN_FILE := cmd/main.go
SERVICE_NAME := consumer
DOCKER_OWNER := phuoctran
IMAGE_VER := v1 

# Debug flags
GCFLAGS := all=-N -l

.PHONY: help build run build-debug debug up down docker-build docker-tag docker-push clean

help: ## Show all available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*##"}; {printf "  \033[32m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the Go application
	@echo "==> Building $(APP_NAME)"
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_FILE)

run: build ## Run the application
	./$(BIN_DIR)/$(APP_NAME)

build-debug: ## Build the app in debug mode
	@echo "==> Building debug binary"
	mkdir -p $(BIN_DIR)
	go build -gcflags="$(GCFLAGS)" -o $(BIN_DIR)/$(APP_NAME)-debug $(MAIN_FILE)

debug: build-debug ## Run the app with delve debugger
	@command -v dlv >/dev/null 2>&1 || go install github.com/go-delve/delve/cmd/dlv@latest
	dlv exec ./$(BIN_DIR)/$(APP_NAME)-debug

clean: ## Clean build artifacts
	rm -rf $(BIN_DIR)
