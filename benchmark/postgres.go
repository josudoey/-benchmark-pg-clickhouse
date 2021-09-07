package benchmark

import (
	"github.com/go-pg/pg/v10"
	"github.com/josudoey/benchmark-pg-clickhouse/model"
)

func PostgresInsert(db *pg.DB, items <-chan *model.PostMeasurement) error {
	for measurement := range items {
		_, err := db.Model(measurement).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}
