# App Layer

## Description

This package contains the dependency injection and initialization logic for the application layers.

## Purpose

The app layer serves as the composition root where all dependencies are wired together following the clean architecture pattern.

## Usage

```go
// Example initialization flow
repositories := app.NewRepository(db)
useCases := app.NewUseCase(repositories)
handlers := app.NewHandler(useCases)
```

## Used By

- Main application entry point
- Server initialization

## Guidelines

- Keep initialization logic centralized
- Follow dependency injection principles
- Maintain clean architecture separation
- Wire dependencies in correct order
