.PHONY: run migrate rollback test

run:
	go run ./cmd/main.go

migrate:
	go run ./cmd/migration.go

rollback:
	go run ./cmd/rollback.go

test:
	go test ./tests/unit/controllerUnitTest