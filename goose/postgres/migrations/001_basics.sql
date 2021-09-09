-- +goose Up
CREATE TABLE IF NOT EXISTS post_measurements (
  id serial PRIMARY KEY,
  member_id bigint NOT NULL,
  post_id bigint NOT NULL,
  type text NOT NULL,
  date date NOT NULL,
  quantity bigint NOT NULL DEFAULT 1,
  created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX post_measurements_member_id_date_idx ON post_measurements(member_id int8_ops,date date_ops);

-- +goose Down
DROP TABLE IF EXISTS post_measurements;
