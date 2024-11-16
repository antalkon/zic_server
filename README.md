# zic_server

 Инструкция - https://principled-crate-062.notion.site/workflow-1404a6405165801fbd61d936d8bee499?pvs=4
## Зависимости:
### Go:
- Верия 1.22!

### Docker:
- **База данных пше** (port: 5444; user: alex; password: alex; db: ZIC-SERVER-main)
```
docker run -d --name ZIC_DB_pg -e POSTGRES_USER=alex -e POSTGRES_PASSWORD=alex -e POSTGRES_DB=ZIC-SERVER-main -p 5444:5432 postgres:latest
```

## Запуск сервера:
Сервер запускатся на порту 8385, localhost.
- Запуск с помощью Makefile
```make run```
- Запуск в ручную
```go run cmd/app/main.go```

## Данные
- Данные от панели:
логин: `admin`
пароль: `admin`

## Project Team:
antalkon - https://github.com/antalkon
