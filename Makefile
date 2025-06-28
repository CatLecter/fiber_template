.PHONY: build run test clean docker-build docker-run

# Build the application
build:
	go build -o bin/app cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -f build/Dockerfile -t fiber-template .

# Run with Docker Compose
docker-run:
	docker compose -f deployments/docker-compose.yml up --build

# Stop Docker Compose
docker-stop:
	docker compose -f deployments/docker-compose.yml down

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate swagger docs
swagger:
	swag init -g cmd/main.go -o docs/api
