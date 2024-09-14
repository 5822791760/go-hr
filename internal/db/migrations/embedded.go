package migrations

import "embed"

//go:embed hr/*.sql
var hrembed embed.FS

func NewHrMigration() (embed.FS, string, string) {
	return hrembed, "hr", "postgres"
}
