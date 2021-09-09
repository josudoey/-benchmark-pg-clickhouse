package driver

import (
	"context"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/josudoey/benchmark-pg-clickhouse/model"
)

func NewPostgresDB(postgresURL string) (*pg.DB, error) {
	opts, err := pg.ParseURL(postgresURL)
	if err != nil {
		return nil, err
	}
	return pg.Connect(opts), nil
}

func NewDefaultPostgresDB() (*pg.DB, error) {
	return NewPostgresDB(os.Getenv("POSTGRES_URL"))
}

func PostgresInsertPostMeasurements(db *pg.DB, items <-chan *model.PostMeasurement) error {
	ctx := context.Background()
	return db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		for measurement := range items {
			_, err := tx.Model(measurement).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
}
