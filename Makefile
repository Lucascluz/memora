.PHONY: build test clean run-server run-examples docker-build docker-run proto-gen fmt

# Build all components
build:
	@echo "Building server..."
	cd server && go build -o memora-server ./cmd
	@echo "Building examples..."
	cd examples && go build -o memora-example ./cmd

# Test all modules
test:
	@echo "Testing server..."
	cd server && go test ./...
	@echo "Testing client..."
	cd client && go test ./...
	@echo "Testing proto..."
	cd proto && go test ./...

# Clean build artifacts
clean:
	rm -f server/memora-server
	rm -f examples/memora-example
	rm -f server/server.log

# Run the server
run-server:
	cd server && go run ./cmd

# Run examples (requires server to be running)
run-examples:
	cd examples && go run ./cmd

# Build Docker image
docker-build:
	docker build -t memora-server .

# Run with Docker
docker-run:
	docker run -p 1212:1212 memora-server

# Run with Docker Compose
docker-up:
	docker-compose up --build

# Generate protobuf files
proto-gen:
	cd proto && PATH=$$PATH:$$HOME/go/bin protoc --go_out=. --go-grpc_out=. memora.proto

# Format all Go files
fmt:
	find . -name "*.go" -exec go fmt {} \;

# Tidy all modules
tidy:
	cd server && go mod tidy
	cd client && go mod tidy
	cd proto && go mod tidy
	cd examples && go mod tidy

# Full integration test
integration-test: build
	@echo "Starting integration test..."
	cd server && nohup ./memora-server > server.log 2>&1 & echo $$! > server.pid
	@sleep 2
	cd examples && ./memora-example
	@if [ -f server/server.pid ]; then kill `cat server/server.pid` && rm server/server.pid; fi
	@echo "Integration test completed successfully!"

# Help
help:
	@echo "Available targets:"
	@echo "  build           - Build all components"
	@echo "  test            - Run tests for all modules"
	@echo "  clean           - Clean build artifacts"
	@echo "  run-server      - Run the server"
	@echo "  run-examples    - Run examples"
	@echo "  docker-build    - Build Docker image"
	@echo "  docker-run      - Run with Docker"
	@echo "  docker-up       - Run with Docker Compose"
	@echo "  proto-gen       - Generate protobuf files"
	@echo "  fmt             - Format all Go files"
	@echo "  tidy            - Tidy all modules"
	@echo "  integration-test - Run full integration test"