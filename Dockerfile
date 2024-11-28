# Этап сборки
FROM golang:1.22 as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o server cmd/app/main.go

# Финальный минимальный, но функциональный образ
FROM debian:bookworm-slim

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем базовые зависимости, корневые сертификаты и инструменты
RUN apt-get update && apt-get install -y --no-install-recommends \
    bash \
    curl \
    ca-certificates \
    net-tools \
    iputils-ping \
    vim \
    libc6 \
    libstdc++6 \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Копируем собранный сервер
COPY --from=builder /app/server .

# Копируем необходимые данные (опционально)
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/data ./data
COPY --from=builder /app/web ./web
COPY --from=builder /app/zic.ico ./zic.ico

# Указываем порт приложения
EXPOSE 8385

# Точка входа
CMD ["./server"]