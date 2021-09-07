package benchmark

import (
	"database/sql"
	"time"

	"github.com/josudoey/benchmark-pg-clickhouse/model"
)

func ClickHouseInsertSamplePostMeasurements(db *sql.DB, items <-chan *model.PostMeasurement) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO post_measurements (id, member_id, post_id, type, date, quantity, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for measurement := range items {
		now := time.Now()
		if _, err := stmt.Exec(
			Rand.Int63(),
			measurement.MemeberID,
			measurement.PostID,
			measurement.Type,
			measurement.Date,
			measurement.Quantity,
			now,
			now,
		); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
