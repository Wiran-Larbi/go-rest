build:
	@go build -o bin/mdes-cs-server cmd/main.go

testing:
	@go test -v ./...

run: build
	@./bin/mdes-cs-server

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down