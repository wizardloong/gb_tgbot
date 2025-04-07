# Stage Building
FROM golang:1.21-alpine AS builder

# Default directory
WORKDIR /app

# Install dependencies
COPY ../source/go.mod ../source/
RUN go mod tidy

# Copy Project files
COPY ../ .

# Устанавливаем переменную окружения (по умолчанию - development)
ARG APP_ENV
ENV APP_ENV=${APP_ENV}

RUN echo "$APP_ENV"

# Билдим бинарник только в продакшене
RUN if [ "$APP_ENV" = "production" ]; then go build -v -o /usr/local/bin/main ./source; fi


# Stage Running
FROM golang:1.21-alpine

# Copy binary from builder
RUN if [ "$APP_ENV" = "production" ]; then cp /usr/local/bin/main /usr/local/bin/main; fi || true

# Check if main binanary exists
RUN ls -la /usr/local/bin/main && \
    chmod +x /usr/local/bin/main && \
    /usr/local/bin/main --version || echo "Binary test fail"

# Set workdir for application
WORKDIR /app

# Устанавливаем переменную окружения (по умолчанию - development)
ARG APP_ENV
ENV APP_ENV=${APP_ENV}
RUN echo "$APP_ENV"

COPY . .

# Run script from /usr/local/bin
CMD ["/bin/sh", "-c", "if [ \"$APP_ENV\" = \"production\" ]; then /usr/local/bin/main; else go run ./source/; fi"]
