# syntax=docker/dockerfile:1

# Build arguments
ARG GO_VERSION=1.25.3
ARG BUILD_ENV=production

# Base stage - common for all environments
FROM golang:${GO_VERSION}-alpine AS base
WORKDIR /app

# Install build dependencies and tools
# gcc, g++, make, musl-dev: Required for CGO and building Go packages with C dependencies
# git: Required for go mod download from git repositories
# ca-certificates: For HTTPS connections
# tzdata: For timezone support
RUN apk add --no-cache \
    gcc \
    g++ \
    make \
    musl-dev \
    git \
    ca-certificates \
    tzdata \
    curl \
    bash

# Copy dependency files
COPY go.mod ./
COPY go.sum ./

# Download dependencies and verify
RUN go mod download && \
    go mod verify

# Development stage
FROM base AS development
ARG AIR_VERSION=latest
ARG MIGRATE_VERSION=v4.18.1

# Install Air for hot-reload and golang-migrate for database migrations
RUN go install github.com/air-verse/air@${AIR_VERSION} && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@${MIGRATE_VERSION}

# Environment variables for development
ENV GO_ENV=development \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Builder stage for production
FROM base AS builder

# Environment variables for building
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy source code
COPY . .

# Build the application with proper path
RUN go build -ldflags="-w -s" -o /app/bin/server ./cmd/main.go

# Production stage
FROM alpine:latest AS production

# Install runtime dependencies
RUN apk --no-cache add \
    ca-certificates \
    tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder --chown=appuser:appuser /app/bin/server /app/server

# Use non-root user
USER appuser

# Environment variables for production
ENV GO_ENV=production \
    GIN_MODE=release

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/app/server", "--health-check"] || exit 1

# Run the application
CMD ["/app/server"]

# Final stage selection based on BUILD_ENV
FROM ${BUILD_ENV} AS final

