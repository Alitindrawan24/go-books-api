# Environment configuration
ENV=$(shell echo ${env} | tr '[:upper:]' '[:lower:]')

ifeq ($(ENV), testing)
	include .env.testing
	USED_ENV=".env.testing"
else
	include .env
	USED_ENV=".env"
endif

# ============================================
# Build & Run
# ============================================

build:
	@go build -v -o go-books-api ./cmd/*.go

run:
	@echo "RUN go-books-api..."
	make build
	@./go-books-api

# ============================================
# Testing
# ============================================

test:
	@go test ./internal/...

# Usage: make test-it env=testing
test-it:
	@echo "Using environment: ${USED_ENV}"
	@go test -tags=integration -race ./it/... -v

test-it-clean:
	@echo "Using environment: ${USED_ENV}"
	@go clean -testcache
	@go test -tags=integration -race ./it/... -v

test-clean:
	@go clean -testcache
	@go test ./internal/...

test-coverage:
	@go test ./internal/... -cover -v

test-coverage-html:
	@go test ./internal/... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

# ============================================
# Mock Generation
# ============================================

# Usage: make mock layer=usecase package=user
mock:
	@mockgen -source=internal/$(layer)/$(package)/$(layer).go -destination=gomock/$(layer)/mock$(package)/mock.go -package=mock$(package)

# Usage: make mock-usecase package=user
mock-usecase:
	@mockgen -source=internal/usecase/$(package)/usecase.go -destination=gomock/usecase/mock$(package)/mock.go -package=mock$(package)

# Usage: make mock-repository-api package=external
mock-repository-api:
	@mockgen -source=internal/repository/api/$(package)/repository.go -destination=gomock/repository/api/mock$(package)/mock.go -package=mock$(package)

# ============================================
# Code Quality
# ============================================

lint:
	@golangci-lint run

# ============================================
# Docker
# ============================================

docker-build:
	@docker-compose build

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

docker-logs:
	@docker-compose logs -f

docker-restart:
	@docker-compose restart

docker-clean:
	@docker-compose down -v