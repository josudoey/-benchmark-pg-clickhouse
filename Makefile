PATH := ${CURDIR}/bin:$(PATH)

.PHONY: goose-postgres-up
goose-postgres-up: bin/goose
	bin/goose -dir=./goose/postgres/migrations postgres '${POSTGRES_URL}' up

.PHONY: goose-postgres-down
goose-postgres-down: bin/goose
	bin/goose -dir=./goose/postgres/migrations postgres '${POSTGRES_URL}' down

.PHONY: goose-clickhouse-up
goose-clickhouse-up: bin/goose
	bin/goose -dir=./goose/clickhouse/migrations clickhouse '${CLICKHOUSE_URL}' up

.PHONY: goose-clickhouse-down
goose-clickhouse-down: bin/goose
	bin/goose -dir=./goose/clickhouse/migrations clickhouse '${CLICKHOUSE_URL}' down


.PHONY: wire
wire: bin/wire
	bin/wire  ./...

bin/wire: go.sum
	GOBIN=$(abspath bin) go install github.com/google/wire/cmd/wire@v0.5.0

bin/goose: go.sum
	GOBIN=$(abspath bin) go install github.com/pressly/goose/v3/cmd/goose@v3.1.0
