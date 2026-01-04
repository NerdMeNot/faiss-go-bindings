# Makefile for faiss-go-bindings
#
# This is a bindings-only module with pre-built static libraries.
# Most users just need: go build / go test

.PHONY: build test clean fmt lint help

# Default target
all: build test

# Build the module
build:
	go build -v ./...

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Clean build artifacts
clean:
	rm -f coverage.out
	go clean

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run --timeout=5m

# Build c_api_ext (requires FAISS headers)
build-ext:
	cd c_api_ext && chmod +x build.sh && ./build.sh

# Rebuild static libraries (runs in CI, requires native platform)
rebuild-libs:
	chmod +x scripts/build-static-libs.sh
	./scripts/build-static-libs.sh

# Show help
help:
	@echo "faiss-go-bindings Makefile"
	@echo ""
	@echo "Targets:"
	@echo "  build          - Build the module"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  clean          - Clean build artifacts"
	@echo "  fmt            - Format code"
	@echo "  lint           - Run golangci-lint"
	@echo "  build-ext      - Build c_api_ext (requires FAISS headers)"
	@echo "  rebuild-libs   - Rebuild static libraries (native platform only)"
	@echo "  help           - Show this help"
