#!make
include .env

go-gen:
	rm -rf internal/mocks/*
	go generate ./...

run-dev:
	go run main.go

test-setup:
	go install github.com/onsi/ginkgo/v2/ginkgo@v2.11.0
	
test: test-setup
	ginkgo ./...
	
test-detail: test-setup
	ginkgo -vv ./...

# go install --tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# use like this -> make migrateinit name=something
migrate:
	migrate create -ext sql -dir pkg/db/migration $(name)
	
migrateup:
	migrate -path pkg/db/migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migratedown:
	migrate -path pkg/db/migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local" -verbose down

.PHONY: migrate migrateup migratedown