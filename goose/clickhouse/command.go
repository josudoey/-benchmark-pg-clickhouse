package clickhouse

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var FS embed.FS

func Up(clickhouseURL string) error {
	goose.SetBaseFS(FS)
	goose.SetDialect("clickhouse")
	db, err := sql.Open("clickhouse", clickhouseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	return goose.Up(db, "migrations")
}

func Down(clickhouseURL string) error {
	goose.SetBaseFS(FS)
	goose.SetDialect("clickhouse")
	db, err := sql.Open("clickhouse", clickhouseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	return goose.Down(db, "migrations")
}
