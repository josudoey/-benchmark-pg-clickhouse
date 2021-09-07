package clickhouse

import (
	"database/sql"
)

func NewDB(clickhouseURL string) (*sql.DB, error) {
	return sql.Open("clickhouse", clickhouseURL)
}
