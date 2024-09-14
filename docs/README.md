# Binaries

```
$ go install github.com/air-verse/air@latest
$ go install github.com/pressly/goose/v3/cmd/goose@latest
$ go install github.com/go-jet/jet/v2/cmd/jet@latest
```

# Running Project

```
$ make dev
```

# Migration

## Migrate
```
$ make db-up
```

## Revert latest
```
$ make db-down
```

## Generate Empty migration
```
$ make db-new
```

# Testing

## Generate Mock Repo
```
$ make mock-irepo
```

## Running Test
```
$ make test-backend
```