package model

import (
	"math/rand"
	"time"
)

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

var Rand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateBenchPostMeasurements(memberID int64) <-chan *PostMeasurement {
	postCount := 100
	dayCount := 365
	typ := "viewed"
	maxQuantity := int64(1000)
	ch := make(chan *PostMeasurement, 1)
	go func() {
		for m := range GenerateMemberPostMeasurements(memberID, postCount, dayCount, typ, maxQuantity) {
			ch <- m
		}
		close(ch)
	}()
	return ch
}

func GenerateMemberPostMeasurements(memberID int64, postCount int, dayCount int, typ string, maxQuantity int64) <-chan *PostMeasurement {
	now := time.Now()
	ch := make(chan *PostMeasurement, 1)
	go func() {
		for p := 0; p < postCount; p++ {
			postID := Rand.Int63()
			for d := 0; d < dayCount; d++ {
				date := now.AddDate(0, 0, -d)
				ch <- &PostMeasurement{
					MemeberID: memberID,
					PostID:    postID,
					Type:      typ,
					Date:      &date,
					Quantity:  Rand.Int63n(maxQuantity) + 1,
				}
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateSamplePostMeasurements(postCount int, memberCount int, dayCount int, typ string, beginPostID int64, beginMemberID int64, maxQuantity int64) <-chan *PostMeasurement {
	postMemberMap := map[int64]int64{}
	postID := beginPostID
	memberID := beginMemberID
	for m := 0; m < memberCount; m++ {
		for p := 0; p < postCount; p++ {
			postMemberMap[postID] = memberID
			postID += 1
		}
		memberID += 1
	}
	now := time.Now()

	ch := make(chan *PostMeasurement, 1)
	go func() {
		for d := 0; d < dayCount; d++ {
			date := now.AddDate(0, 0, -d)
			for postID, memberID := range postMemberMap {
				ch <- &PostMeasurement{
					MemeberID: memberID,
					PostID:    postID,
					Type:      typ,
					Date:      &date,
					Quantity:  Rand.Int63n(maxQuantity) + 1,
				}
			}
		}
		close(ch)
	}()
	return ch
}
