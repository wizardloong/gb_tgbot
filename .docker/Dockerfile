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

# Устанавливаем Git и зависимости
# RUN apk add --no-cache git

# # Проверяем версию Git (опционально)
# RUN git --version

# Copy binary from builder
RUN if [ "$APP_ENV" = "production" ]; then cp /usr/local/bin/main /usr/local/bin/main; fi || true

# Установка зависимостей для MySQL (если нужны)
# RUN apk add --no-cache gcc musl-dev git

#RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Устанавливаем curl и распаковываем migrate
RUN apk add --no-cache curl tar \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin


# Добавление бинарника migrate в PATH (опционально)
ENV PATH=$PATH:/go/bin


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
CMD ["/bin/sh", "-c", "if [ \"$APP_ENV\" = \"production\" ]; then /usr/local/bin/main; else cd source && go run .; fi"]
