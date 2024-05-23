LOCAL_BIN:=$(CURDIR)/bin

prepare:
	go mod download
	make migrate

docs:
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@1.16.3
	swag init -g ./internal/app/app.go


run-postgres:
	docker run --name library-postgres -e POSTGRES_PASSWORD=qwerty123456 -dp 5432:5432 postgres:16

migrate:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.20.0
	GOBIN=$(LOCAL_BIN) goose -dir migrations postgres "host=localhost port=5432 dbname=postgres user=postgres password=qwerty123456 sslmode=disable" up -v
