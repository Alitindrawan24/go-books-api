# Constants

## Description

This package contains application-wide constants and enumerations.

## Purpose

Centralized location for all constant values used throughout the application to ensure consistency and maintainability.

## Usage

```go
import "github.com/Alitindrawan24/go-books-api/internal/constant"

// Using header constants
clientIP := r.Header.Get(string(constant.ClientIPAddress))
```

## Used By

- Handler layer
- Middleware layer
- Use case layer
- Repository layer

## Guidelines

- Group related constants together
- Use descriptive constant names
- Avoid magic numbers and strings
- Document complex constant values
