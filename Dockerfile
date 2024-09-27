# Используем официальный образ Go в качестве базового
FROM golang:1.22.2 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код приложения
COPY cmd/myapp .

# Собираем приложение
RUN go build -o myapp ./cmd/myapp

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk --no-cache add ca-certificates

# Копируем собранное приложение из образа builder
COPY --from=builder /app/myapp .

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

# Команда для запуска приложения
CMD ["./myapp"]
