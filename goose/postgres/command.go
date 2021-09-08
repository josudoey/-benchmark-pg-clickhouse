package postgres

import (
	"database/sql"
	"embed"

	"github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var FS embed.FS

func NewDB(postgresURL string) (*sql.DB, error) {
	connector, err := pq.NewConnector(postgresURL)
	if err != nil {
		return nil, err
	}
	return sql.OpenDB(connector), nil
}

func Up(postgresURL string) error {
	goose.SetBaseFS(FS)
	goose.SetDialect("postgres")
	db, err := NewDB(postgresURL)
	if err != nil {
		return err
	}
	defer db.Close()

	return goose.Up(db, "migrations")
}

func Down(postgresURL string) error {
	goose.SetBaseFS(FS)
	goose.SetDialect("postgres")
	db, err := NewDB(postgresURL)
	if err != nil {
		return err
	}
	defer db.Close()

	return goose.Down(db, "migrations")
}
