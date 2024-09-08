ifneq (,$(wildcard ./.env))
    include .env
    export POSTGRES_USER POSTGRES_PASSWORD POSTGRES_HOST POSTGRES_PORT POSTGRES_DB
endif

DB_STRING="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

.wait-for-pg:
	./internal/scripts/wait-for-postgres.sh

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

dev: up .wait-for-pg
	./internal/scripts/air -c .air.toml

build:
	go build -o ./internal/scripts ./cmd/app

start:
	./internal/scripts/api

gen:
	./internal/scripts/jet -dsn=${DB_STRING} -schema=public -path=./internal/db

drop-db:
	docker-compose up -d postgres
	make .wait-for-pg
	docker-compose exec postgres dropdb -U ${DB_USER} --if-exists ${DB_DATABASE}
	docker-compose exec postgres createdb -U ${DB_USER} ${DB_DATABASE}

reset-db: drop-db db-up gen

db-status:
	./internal/scripts/goose -dir="./internal/db/postgres/migrations" postgres ${DB_STRING} status

db-up:
	./internal/scripts/goose -dir="./internal/db/postgres/migrations" postgres ${DB_STRING} up

db-down:
	./internal/scripts/goose -dir="./internal/db/postgres/migrations" postgres ${DB_STRING} down

db-redo:
	./internal/scripts/goose -dir="./internal/db/postgres/migrations" postgres ${DB_STRING} redo

db-new:
	./internal/scripts/goose -dir="./internal/db/postgres/migrations" postgres ${DB_STRING} create ${name} sql
