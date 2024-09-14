package dbutil

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

type migrationFunc func() (embed.FS, string, string)

func AutoMigrate(db *sql.DB, migrationFunc migrationFunc) error {
	file, dir, dialect := migrationFunc()

	goose.SetBaseFS(file)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	if err := goose.Up(db, dir); err != nil {
		return err
	}

	return nil
}
