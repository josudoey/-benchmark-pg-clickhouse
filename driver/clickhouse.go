package driver

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/josudoey/benchmark-pg-clickhouse/model"
)

func NewClickHouseDB(clickhouseURL string) (*sql.DB, error) {
	return sql.Open("clickhouse", clickhouseURL)
}

func NewDefaultClickHouseDB() (*sql.DB, error) {
	return NewClickHouseDB(os.Getenv("CLICKHOUSE_URL"))
}

func ClickHouseInsertPostMeasurements(db *sql.DB, items <-chan *model.PostMeasurement) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO post_measurements (id, member_id, post_id, type, date, quantity, created_at, updated_at)
VALUES (@ID, @MemeberID, @PostID, @Type, @Date, @Quantity, @CreatedAt, @UpdatedAt)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for measurement := range items {
		now := time.Now()
		if _, err := stmt.Exec(
			sql.Named("MemeberID", measurement.MemeberID),
			sql.Named("ID", model.Rand.Int63()),
			sql.Named("PostID", measurement.PostID),
			sql.Named("Type", measurement.Type),
			sql.Named("Date", measurement.Date),
			sql.Named("Quantity", measurement.Quantity),
			sql.Named("CreatedAt", now),
			sql.Named("UpdatedAt", now),
		); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
