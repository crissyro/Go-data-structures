MODULE := github.com/crissyro/go-data-structures
COVERAGE_FILE := coverage.out
BIN_DIR := bin
GOTESTSUM := gotest.tools/gotestsum

.PHONY: all setup test coverage report lint build clean

all: test build

setup:
	@echo "⚙️  Installing tools..."
	go install $(GOTESTSUM)@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test:
	@echo "🧪 Running tests..."
	$(GOPATH)/bin/gotestsum --format dots -- -cover -v ./...

coverage:
	@echo "📊 Generating coverage..."
	go test -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...

report: coverage
	@echo "📈 Opening coverage report..."
	go tool cover -html=$(COVERAGE_FILE)

lint:
	@echo "🔍 Linting code..."
	$(GOPATH)/bin/golangci-lint run ./... --timeout 5m

build:
	@echo "🔨 Building binaries..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/ ./...

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(COVERAGE_FILE) coverage.html $(BIN_DIR)