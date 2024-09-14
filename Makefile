ifneq (,$(wildcard ./.env))
    include .env
    export POSTGRES_USER POSTGRES_PASSWORD POSTGRES_HOST POSTGRES_PORT POSTGRES_DB
endif

DB_STRING="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"
MIGRATION_DIR="./internal/db/migrations"

.wait-for-pg:
	./scripts/wait-for-postgres.sh
	

.install-depen:
	go install github.com/air-verse/air@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/go-jet/jet/v2/cmd/jet@latest

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

dev: up .wait-for-pg
	./scripts/air -c .air.toml

build:
	go build -o ./scripts ./cmd/backend

start:
	./scripts/api

gen:
	./scripts/jet -dsn=${DB_STRING} -path=./internal/db/schema

drop-db:
	docker-compose up -d postgres
	make .wait-for-pg
	docker-compose exec postgres dropdb -U ${DB_USER} --if-exists ${DB_DATABASE}
	docker-compose exec postgres createdb -U ${DB_USER} ${DB_DATABASE}

reset-db: drop-db db-up gen

db-status:
	./scripts/goose -dir=${MIGRATION_DIR} postgres ${DB_STRING} status

db-up:
	./scripts/goose -dir=${MIGRATION_DIR} postgres ${DB_STRING} up

db-down:
	./scripts/goose -dir=${MIGRATION_DIR} postgres ${DB_STRING} down

db-redo:
	./scripts/goose -dir=${MIGRATION_DIR} postgres ${DB_STRING} redo

db-new:
	./scripts/goose -dir=${MIGRATION_DIR} postgres ${DB_STRING} create ${name} sql


mock-irepo:
	@find internal/repos -name '*.go' ! -name 'entity.go' | while read file; do \
		dirname=$$(dirname $$file); \
		basefile=$$(basename $$file); \
		mockpath=test/mocks/repos/mock_$$(basename $$dirname)/mock_$$basefile; \
		mockgen -source=$$file -destination=$$mockpath -package=mock_$$(basename mock_$$dirname); \
	done

test-backend:
	go test ./test/backend/...

doc:
	./scripts/godoc -http :8080