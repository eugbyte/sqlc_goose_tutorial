.PHONY: all test clean	# https://stackoverflow.com/a/3931814/6514532

start:
	go run ./cmd/app

create-migration:
	goose -dir=internal/db/migration -s create create_book sql

apply-migration:
	goose -dir=internal/db/migration postgres "user=postgres password=postgres dbname=book_demo sslmode=disable" up

rollback-migration-to:
	goose -dir=internal/db/migration postgres "user=postgres password=postgres dbname=book_demo sslmode=disable" down-to 123456

sqlc:
	sqlc generate