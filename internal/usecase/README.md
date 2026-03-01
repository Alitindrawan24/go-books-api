# Use Case

## Description

This package contains business logic and application use cases following clean architecture principles.

## Purpose

Implements the core business logic and orchestrates interactions between different layers of the application.

## Usage

```go
// Use case initialization
tokenUseCase := access_token.New(repository, config)

// Handler usage
result, err := tokenUseCase.GenerateToken(ctx, request)
```

## Used By

- Handler layer

## Guidelines

- Keep use cases focused on business logic only
- Avoid direct dependencies on external frameworks
- Use interfaces for repository dependencies
- Handle business errors appropriately
- Make use cases testable with clear inputs/outputs
