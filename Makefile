run:
	go run ./cmd/main.go

migrate:
	go run ./migrations/*.go

.PHONY: run migrate
