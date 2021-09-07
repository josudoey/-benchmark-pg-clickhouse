package postgres

import (
	"database/sql"

	"github.com/lib/pq"
)

func NewDB(postgresURL string) (*sql.DB, error) {
	connector, err := pq.NewConnector(postgresURL)
	if err != nil {
		return nil, err
	}
	return sql.OpenDB(connector), nil
}
