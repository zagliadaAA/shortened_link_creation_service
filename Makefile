migrate-up:
	goose -dir ./db/migrations postgres "host=127.0.0.1 port=5433 user=postgres password=123 dbname=shortened_link_creation_service sslmode=disable" up

migrate-create:
	goose -dir db/migrations create $(F) sql

lint:
	golangci-lint run ./...