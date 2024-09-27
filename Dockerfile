# Используем официальный образ Golang как базовый
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Собираем приложение
RUN go build -o myapp ./cmd/myapp

# Используем легковесный образ для финальной сборки
FROM alpine:latest
WORKDIR /root/

# Копируем скомпилированное приложение из образа builder
COPY --from=builder /app/myapp .

# Указываем команду для запуска приложения
CMD ["./myapp"]
