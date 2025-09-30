# Build stage
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy all module files
COPY . .

# Download dependencies for server
WORKDIR /app/server
RUN go mod download

# Build the server application
RUN CGO_ENABLED=0 GOOS=linux go build -o memora-server ./cmd

# Runtime stage
FROM alpine:latest

# Install ca-certificates for SSL/TLS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 memora && \
    adduser -u 1001 -G memora -s /bin/sh -D memora

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server/memora-server .

# Change ownership to non-root user
RUN chown memora:memora /app/memora-server

# Switch to non-root user
USER memora

# Expose port
EXPOSE 1212

# Health check (simple check that process is running)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD pgrep memora-server || exit 1

# Run the binary
CMD ["./memora-server"]