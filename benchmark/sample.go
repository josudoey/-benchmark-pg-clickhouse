package benchmark

import (
	"math/rand"
	"time"

	"github.com/josudoey/benchmark-pg-clickhouse/model"
)

var Rand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateDefaultPostMeasurements() <-chan *model.PostMeasurement {
	postCount := 100
	memberCount := 1000
	dayCount := 365
	typ := "viewed"
	beginPostID := int64(1)
	beginMemberID := int64(1)
	maxQuantity := int64(1000)
	return GenerateSamplePostMeasurements(postCount, memberCount, dayCount, typ, beginPostID, beginMemberID, maxQuantity)
}

func GenerateSamplePostMeasurements(postCount int, memberCount int, dayCount int, typ string, beginPostID int64, beginMemberID int64, maxQuantity int64) <-chan *model.PostMeasurement {
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

	ch := make(chan *model.PostMeasurement, 1)
	go func() {
		for d := 0; d < dayCount; d++ {
			date := now.AddDate(0, 0, -d)
			for postID, memberID := range postMemberMap {
				ch <- &model.PostMeasurement{
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
