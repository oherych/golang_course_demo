package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func Run(connString string) error {
	goose.SetBaseFS(embedMigrations)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("connect to DB: %w", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error setting dialect: %w", err)
	}

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}
