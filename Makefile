PATH := ${CURDIR}/bin:$(PATH)


.PHONY: bench
bench: bench-clickhouse bench-postgres

.PHONY: clean
clean: goose-postgres-up goose-postgres-down goose-clickhouse-up goose-clickhouse-down

.PHONY: bench-clickhouse
bench-clickhouse: goose-clickhouse-up bench-clickhouse-insert bench-clickhouse-query

.PHONY: bench-clickhouse-insert
bench-clickhouse-insert:
	go test -bench=BenchmarkClickHouseInsert -benchtime=1x -run=None -benchmem  ./driver/...

.PHONY: bench-clickhouse-query
bench-clickhouse-query:
	go test -bench=BenchmarkClickHouseQuery -benchtime=1x -run=None -benchmem  ./driver/...

.PHONY: bench-postgres
bench-postgres: goose-postgres-up bench-postgres-insert bench-postgres-query

.PHONY: bench-postgres-insert
bench-postgres-insert:
	go test -bench=BenchmarkPostgresInsert -benchtime=1x -run=None -benchmem  -timeout 2h ./driver/...

.PHONY: bench-postgres-query
bench-postgres-query:
	go test -bench=BenchmarkPostgresQuery -benchtime=1x -run=None -benchmem  ./driver/...

.PHONY: goose-postgres-up
goose-postgres-up: bin/goose
	@bin/goose -dir=./goose/postgres/migrations postgres '${POSTGRES_URL}' up

.PHONY: goose-clickhouse-up
goose-clickhouse-up: bin/goose
	@bin/goose -dir=./goose/clickhouse/migrations clickhouse '${CLICKHOUSE_URL}' up

.PHONY: goose-postgres-down
goose-postgres-down: bin/goose
	@bin/goose -dir=./goose/postgres/migrations postgres '${POSTGRES_URL}' down

.PHONY: goose-clickhouse-down
goose-clickhouse-down: bin/goose
	@bin/goose -dir=./goose/clickhouse/migrations clickhouse '${CLICKHOUSE_URL}' down

.PHONY: wire
wire: bin/wire
	bin/wire  ./...

bin/wire: go.sum
	GOBIN=$(abspath bin) go install github.com/google/wire/cmd/wire@v0.5.0

bin/goose: go.sum
	GOBIN=$(abspath bin) go install github.com/pressly/goose/v3/cmd/goose@v3.1.0
