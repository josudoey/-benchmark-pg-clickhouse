package main

import (
	"log"
	"os"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/josudoey/benchmark-pg-clickhouse/driver"
	goosech "github.com/josudoey/benchmark-pg-clickhouse/goose/clickhouse"
	"github.com/josudoey/benchmark-pg-clickhouse/model"
	"github.com/spf13/cobra"
)

type CommandOptions struct {
	clickhouseURL string
	postCount     int
	memberCount   int
	dayCount      int
	typ           string
	beginPostID   int64
	beginMemberID int64
	maxQuantity   int64
}

func NewCommandOptions() *CommandOptions {
	return &CommandOptions{
		clickhouseURL: "tcp://127.0.0.1:9000?username=default&password=&database=clicks",
	}
}

func (o *CommandOptions) Run(cmd *cobra.Command, args []string) {
	err := goosech.Up(o.clickhouseURL)
	if err != nil {
		log.Fatal(err)
	}

	db, err := driver.NewClickHouseDB(o.clickhouseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = driver.ClickHouseInsertPostMeasurements(db, model.GenerateSamplePostMeasurements(
		o.postCount,
		o.memberCount,
		o.dayCount,
		o.typ,
		o.beginPostID,
		o.beginMemberID,
		o.maxQuantity,
	))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	o := NewCommandOptions()
	cmd := &cobra.Command{
		Use: "ch-gen-sample",
		Run: o.Run,
	}

	cmd.Flags().StringVar(&o.clickhouseURL, "clickhouse-url", os.Getenv("CLICKHOUSE_URL"), "clickhouse url (default: CLICKHOUSE_URL)")
	cmd.Flags().IntVar(&o.postCount, "post-count", 100, "post count (default: 100)")
	cmd.Flags().IntVar(&o.memberCount, "member-count", 1000, "number of member count for sample (default: 1000)")
	cmd.Flags().IntVar(&o.dayCount, "day-count", 365, "number of day count for sample (default: 365)")
	cmd.Flags().StringVar(&o.typ, "type", "viewed", "type for sample (default: viewed)")
	cmd.Flags().Int64Var(&o.beginPostID, "begin-post-id", 1, "begin id of post for sample (default: 1)")
	cmd.Flags().Int64Var(&o.beginMemberID, "begin-member-id", 1, "begin id of member for sample (default: 1)")
	cmd.Flags().Int64Var(&o.maxQuantity, "max-quantity", 1000, "max quantity for sample (default: 1000)")
	cmd.Execute()
}
