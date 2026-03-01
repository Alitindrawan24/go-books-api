# Go Books API

A Go REST API following clean architecture principles with dependency injection and testability.

## Project Structure

```
.
├── cmd/                    # Application entry point
├── internal/              
│   ├── app/              # Dependency injection
│   ├── constant/         # Application constants
│   ├── entity/           # Domain entities
│   ├── handler/          # HTTP handlers
│   ├── middleware/       # HTTP middleware
│   ├── pkg/              # Shared utilities (config, logger, clocker)
│   ├── repository/       # Data access (mysql, api)
│   ├── schema/           # Request/response DTOs
│   ├── server/           # HTTP server setup
│   └── usecase/          # Business logic
├── gomock/               # Generated mocks
└── storage/              # Storage directory
```

## Installation

1. Clone and navigate to the project:
```bash
git clone https://github.com/Alitindrawan24/go-books-api.git
cd go-books-api
```

2. Copy environment configuration:
```bash
cp .env.example .env
```

3. Install dependencies:
```bash
go mod download
```

4. Run the application:
```bash
make run
```

## Useful Commands

### Build & Run
```bash
make build              # Build binary
make run                # Build and run
```

### Docker
```bash
make docker-up          # Start containers
make docker-down        # Stop containers
make docker-logs        # View logs
make docker-restart     # Restart containers
make docker-clean       # Remove volumes
```

### Mock Generation
```bash
make mock layer=usecase package=user              # General mock
make mock-usecase package=user                    # Use case mock
make mock-repository-api package=external         # API repository mock
```

### Code Quality
```bash
make lint               # Run linter
```

## Testing

```bash
make test                           # Run unit tests
make test-it env=testing            # Run integration tests
make test-coverage                  # Run tests with coverage
make test-coverage-html             # Generate HTML coverage report
make test-clean                     # Clean test cache and run tests
```
