.PHONY: run migrate rollback test watch

run:
	go run ./cmd/main.go

migrate:
	go run ./cmd/migration.go

rollback:
	go run ./cmd/rollback.go

test:
	go test ./tests/unit/controllerUnitTest

watch:
	@if docker ps -f "name=my-postgres-container" --format "{{.Names}}" | grep -q "my-postgres-container"; then \
		echo "PostgreSQL container is already running"; \
	else \
		docker pull postgres; \
		docker run -d --name my-postgres-container -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 postgres; \
	fi
	@if ! psql "postgres://postgres:mysecretpassword@localhost:5432/postgres" -c "SELECT 1 FROM migrations LIMIT 1;" | grep -q "1 row"; then \
		echo "No migrations found. Applying migrations..."; \
		go run ./cmd/migration.go; \
	fi
	go run ./cmd/main.go
