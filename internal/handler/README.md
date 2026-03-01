# Handler

## Description

This package contains HTTP request handlers following the clean architecture pattern.

## Purpose

Handles HTTP requests, validates input, calls appropriate use cases, and returns HTTP responses.

## Usage

```go
// Handler initialization
handler := accesstoken.New(useCase.accessToken, useCase.simpkopdesToken)

// Route registration
router.POST("/api/token", handler.GetAccessToken)
```

## Used By

- Server layer (route registration)

## Guidelines

- Keep handlers thin - delegate business logic to use cases
- Validate all input data
- Return appropriate HTTP status codes
- Handle errors gracefully
