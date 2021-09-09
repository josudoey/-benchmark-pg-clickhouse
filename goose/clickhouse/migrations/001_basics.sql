-- +goose Up
CREATE TABLE post_measurements (
  id Int64,
  member_id Int64,
  post_id Int64,
  type String,
  date Date,
  quantity Int64,
  created_at DateTime64,
  updated_at DateTime64
)
ENGINE = MergeTree()
PARTITION BY toYYYYMM(date)
ORDER BY (id, date, intHash32(member_id))
SAMPLE BY intHash32(member_id);

-- +goose Down
DROP TABLE post_measurements;
