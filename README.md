# zic_server

## Зависимости:
### Go:
- Верия 1.22!

### Docker:
- **База данных Postgres** (port: 5444; user: alex; password: alex; db: ZIC-SERVER-main)
```
docker run -d --name ZIC_DB_pg -e POSTGRES_USER=alex -e POSTGRES_PASSWORD=alex -e POSTGRES_DB=ZIC-SERVER-main -p 5444:5432 postgres:latest
```

## Запуск сервера:
Сервер запускатся на порту 8385, localhost.
- Запуск с помощью Makefile
```make run```
- Запуск в ручную
```go run cmd/app/main.go```

