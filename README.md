# Fiber Template

A Go web application template using Fiber framework with PostgreSQL database.

## Project Structure

This project follows the [Go project layout standards](https://github.com/golang-standards/project-layout).

```
├── cmd/                    # Main applications
├── configs/               # Configuration files
├── docs/                  # Documentation and API specs
├── internal/              # Private application and library code
│   ├── database/         # Database related code
│   ├── handlers/         # HTTP handlers
│   ├── repositories/     # Data access layer
│   ├── schemes/          # Data structures
│   ├── services/         # Business logic
│   └── utils/            # Utility functions
├── build/                # Build and packaging
├── deploy/               # Deployment configurations
├── scripts/              # Scripts for various tasks
└── test/                 # Additional test files
```

## Quick Start

### Using Makefile (Recommended)

```shell
# Copy environment file
cp ./build/example.env ./build/.env

# Run with Docker Compose
make docker-run

# Or run locally
make run
```

### Manual Setup

```shell
# Copy environment file
cp ./build/example.env ./build/.env

# Run with Docker Compose
docker compose -f ./deploy/docker-compose.yml up -d --build
```

## API Documentation

Once the application is running, you can access:

- [Swagger UI](http://localhost:8000/docs)
- API Base URL: `http://localhost:8000/api/v1`

## Available Commands

- `make build` - Build the application
- `make run` - Run the application locally
- `make test` - Run tests
- `make docker-run` - Run with Docker Compose
- `make docker-stop` - Stop Docker Compose
- `make fmt` - Format code
- `make lint` - Lint code
- `make clean` - Clean build artifacts
