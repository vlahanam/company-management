# syntax=docker/dockerfile:1

# Build arguments
ARG GO_VERSION=1.21
ARG BUILD_ENV=production

# Base stage - common for all environments
FROM golang:${GO_VERSION}-alpine AS base
WORKDIR /app

# Install basic dependencies
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata

# Copy dependency files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Development stage
FROM base AS development
ARG AIR_VERSION=latest

# Install Air for hot-reload
RUN go install github.com/air-verse/air@${AIR_VERSION}

# Environment variables for development
ENV GO_ENV=development \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Use Air for hot-reload in development
CMD ["air", "-c", ".air.toml"]

# Builder stage for production
FROM base AS builder

# Environment variables for building
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy source code
COPY . .

# Build the application
RUN go build -ldflags="-w -s" -o /app/bin/server ./cmd/server

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

