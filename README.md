# Zentas Informatics Class - SERVER MVP
About project: http:// .pdf
! MVP VERSION
## INSTALLATION
**USE DOCKER:**
1. Install Docker
2. Run database container
```
docker run -d --name ZIC_DB_pg -e POSTGRES_USER=YOUR_NAME -e POSTGRES_PASSWORD=YOUT_PASSWORD -e POSTGRES_DB=ZIC-SERVER-main -p 5444:5432 postgres:latest
```
(заменить поля YOUR_NAME и YOUR_PASSWORD на подходящие)
3. Run zic container


## Зависимости:
### Go:
- Верия 1.22!

### Docker:
- **База данных postgres** (port: 5444; user: alex; password: alex; db: ZIC-SERVER-main)
```

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
