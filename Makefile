# Makefile for dbcp-smgt

# Directories
BIN_DIR    := bin
CMD_DIR    := ./cmd/dbcp-smgt
PLUGIN_DIR := ./plugins
ETC_DIR    := etc
TESTS_DIR  := ./tests

# Application variables
APP_NAME     := dbcp-smgt
DOCKER_IMAGE := yourorg/$(APP_NAME)  # Change as needed

# Plugin targets (list names of plugins, if any; for dynamic loading)
PLUGINS := postgresql

# Default target: build everything.
.PHONY: all
all: build-plugins build-main

# Build the main binary and place it in bin/
.PHONY: build-main
build-main:
	@echo "Building main binary..."
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-s -w" -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

# Build plugin modules as shared objects and place them in bin/
.PHONY: build-plugins
build-plugins: $(addprefix $(BIN_DIR)/, $(addsuffix .so, $(PLUGINS)))

$(BIN_DIR)/%.so: $(PLUGIN_DIR)/%
	@echo "Building plugin $*..."
	@mkdir -p $(BIN_DIR)
	go build -buildmode=plugin -o $(BIN_DIR)/$*.so $(PLUGIN_DIR)/$*

# Build all
.PHONY: build
build: build-main build-plugins

# Test target: Build dependencies and then run tests.
# Set CONFIG_DIR so tests can use configuration from etc/
.PHONY: test
test: build
	@echo "Running unit tests..."
	CONFIG_DIR=$(ETC_DIR) go test -v $(TESTS_DIR)/unit/...
	@echo "Running integration tests..."
	@if [ -d "$(TESTS_DIR)/integration" ]; then \
		CONFIG_DIR=$(ETC_DIR) go test -v $(TESTS_DIR)/integration/...; \
	else \
		echo "No integration tests found."; \
	fi

# Run the main application.
.PHONY: run
run: build-main
	@echo "Running $(APP_NAME)..."
	./$(BIN_DIR)/$(APP_NAME) --config=$(ETC_DIR)/config.yaml

# Lint the code with golangci-lint.
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run

# Run go vet.
.PHONY: vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Clean build artifacts.
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	go clean
	rm -rf $(BIN_DIR)

# Deploy: Build everything, then build/push a Docker image.
.PHONY: deploy
deploy: build
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):latest .
	@echo "Pushing Docker image..."
	docker push $(DOCKER_IMAGE):latest
