# Stage Building
FROM golang:1.21-alpine AS builder

# Default directory
WORKDIR /app

# Install dependencies
COPY ../go.mod .
RUN go mod download

# Copy Project files
COPY ../ .

# Compile executable
RUN go build -v -o /usr/local/bin/main .

# Stage Running
FROM alpine:latest

# Copy binary from builder
COPY --from=builder /usr/local/bin/main /usr/local/bin/main

# Check if main binanary exists
RUN ls -la /usr/local/bin/main && \
    chmod +x /usr/local/bin/main && \
    /usr/local/bin/main --version || echo "Binary test fail"

# Set workdir for application
WORKDIR /app

# Run script from /usr/local/bin
CMD ["/bin/sh", "-c", "cd /app && /usr/local/bin/main"]