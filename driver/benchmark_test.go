package driver

import (
	"testing"
	"time"

	"github.com/josudoey/benchmark-pg-clickhouse/model"
)

func BenchmarkPostgresInsert(b *testing.B) {
	db, err := NewDefaultPostgresDB()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		err = PostgresInsertPostMeasurements(db, model.GenerateDefaultPostMeasurements())
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPostgresQuery(b *testing.B) {
	db, err := NewDefaultPostgresDB()
	if err != nil {
		b.Fatal(err)
	}
	now := time.Now()
	begin := now.AddDate(0, 0, -365).Format("2006-01") + "-01"
	end := now.Format("2006-01") + "-01"
	for i := 0; i < b.N; i++ {
		_, err := db.Exec(`
SELECT 
  member_id,type, 
  sum(quantity) AS sum 
FROM post_measurements 
WHERE member_id = ? 
  AND (date BETWEEN ? AND ?) 
GROUP BY member_id,type;`,
			i+1, begin, end,
		)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClickHouseInsert(b *testing.B) {
	db, err := NewDefaultClickHouseDB()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		err = ClickHouseInsertPostMeasurements(db, model.GenerateDefaultPostMeasurements())
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClickHouseQuery(b *testing.B) {
	db, err := NewDefaultClickHouseDB()
	if err != nil {
		b.Fatal(err)
	}
	now := time.Now()
	begin := now.AddDate(0, 0, -365)
	end := now
	for i := 0; i < b.N; i++ {
		rows, err := db.Query(`
SELECT 
	member_id,type, 
	sum(quantity) AS sum 
FROM post_measurements 
WHERE member_id = ? 
	AND (toYYYYMM(date) BETWEEN toYYYYMM(?) AND toYYYYMM(?)) 
	GROUP BY member_id,type;`,
			i+1, begin, end,
		)
		if err != nil {
			b.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var (
				memberID int64
				typ      string
				sum      int64
			)
			if err := rows.Scan(&memberID, &typ, &sum); err != nil {
				b.Fatal(err)
			}
		}

		if err := rows.Err(); err != nil {
			b.Fatal(err)
		}
	}
}
