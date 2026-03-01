# Package (pkg)

## Description

This package contains reusable utility packages and shared components.

## Purpose

Houses common utilities, helpers, and shared functionality that can be used across different layers of the application.

## Usage

```go
import (
    "github.com/Alitindrawan24/go-books-api/internal/pkg/config"
    "github.com/Alitindrawan24/go-books-api/internal/pkg/logger"
)

// Load configuration
cfg := config.Load()

// Initialize logger
log := logger.New(cfg.LogLevel)
```

## Used By

- Handler layer
- Middleware layer
- Use case layer
- Repository layer
- Server layer

## Guidelines

- Keep packages focused and cohesive
- Avoid circular dependencies
- Make utilities reusable and testable
- Document public APIs clearly
