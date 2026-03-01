# Schema

## Description

This package contains data schemas, validation rules, and data transfer objects (DTOs).

## Purpose

Defines the structure of data exchanged between different layers and external systems, including validation rules and serialization formats.

## Usage

```go
// Example request schema
type CreateUserRequest struct {
    Name  string `json:"name" validate:"required,min=2"`
    Email string `json:"email" validate:"required,email"`
}

// Example response schema
type UserResponse struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

## Used By

- Handler layer
- Use case layer
- Repository layer

## Guidelines

- Use struct tags for JSON serialization and validation
- Keep schemas focused on data structure only
- Separate request/response schemas from domain entities
- Include comprehensive validation rules
- Document schema fields clearly
