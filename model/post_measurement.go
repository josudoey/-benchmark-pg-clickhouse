package model

import "time"

type PostMeasurement struct {
	ID        int64      `pg:"id,notnull"`
	MemeberID int64      `pg:"member_id,notnull"`
	PostID    int64      `pg:"post_id,notnull"`
	Type      string     `pg:"type,notnull"`
	Date      *time.Time `pg:"date,notnull"`
	Quantity  int64      `pg:"quantity,notnull"`
	CreatedAt *time.Time `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}
