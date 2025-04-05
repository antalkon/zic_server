run-dev:
	go run cmd/backend/main.go

test-run:
	go test -v ./...

upd-swg:
	rm -rf docs
	swag init -g cmd/backend/main.go

dc-l:
	docker-compose -f docker-compose.local.yaml up -d