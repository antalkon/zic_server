# Установка зависимостей
install:
    go mod tidy

# Запуск тестов
test:
    go test ./... -v

# Линтинг
lint:
    golangci-lint run

# Сборка приложения
build:
    go build -o main ./cmd/myapp/main.go

# Локальный запуск приложения
run:
    ./main
