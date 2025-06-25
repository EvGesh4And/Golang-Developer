LOCAL_BIN:=$(CURDIR)/bin

install:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: run_migrations
run_migrations:
	goose -dir db/migrations postgres "postgres://postgres:changeme@localhost:5432/postgres?sslmode=disable" up

.PHONY: status_migrations
status_migrations:
	goose -dir db/migrations postgres "postgres://postgres:changeme@localhost:5432/postgres?sslmode=disable" status


.PHONY: db_connect
db_connect:
	psql -d "host=localhost port=5432 dbname=postgres user=postgres password=changeme"

.PHONY: run_test
run_service:
	go test ./...