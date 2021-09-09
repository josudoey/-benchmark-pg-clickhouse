# benchmark postgres clickhouse

## Summary
|              |             |  clickhouse |    postgres   |
|--------------|-------------|:-----------:|:-------------:|
|              | total count |  36,500,000 |    36,500,000 |
|              | disk usage  | 560,734,742 | 4,589,027,328 |
| insert 1000x | time        |    141.708s |     1864.889s |
|              | ns/op       | 141,200,907 | 1,864,550,944 |
|              | B/op        |  35,162,619 |    41,084,100 |
|              | allocs/op   |     738,508 |       693,606 |
|  query 1000x | time        |     28.779s |       25.785s |
|              | ns/op       |  28,456,499 |    25,294,512 |
|              | B/op        |       8,345 |         1,250 |
|              | allocs/op   |          75 |             4 |

## Table Schema
### Clickhouse 
```sql
CREATE TABLE post_measurements(
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
```


## Postgres
```sql
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
```

## Benchmark Result

### bench-clickhouse-100

```
$ make bench-clickhouse-100
2021/09/09 09:44:23 OK    001_basics.sql
2021/09/09 09:44:23 goose: no migrations to run. current version: 1
go test -bench=BenchmarkClickHouseInsert -benchtime=100x -run=None -benchmem -timeout 1h ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkClickHouseInsert-8          100          83611050 ns/op        35709481 B/op     793825 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      8.486s
go test -bench=BenchmarkClickHouseQuery -benchtime=100x -run=None -benchmem -timeout 1h ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkClickHouseQuery-8           100          10267872 ns/op           46408 B/op         76 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      1.181s
```

### bench-postgres-100

```
$ make bench-postgres-100
2021/09/09 09:50:41 OK    001_basics.sql
2021/09/09 09:50:41 goose: no migrations to run. current version: 1
go test -bench=BenchmarkPostgresInsert -benchtime=100x -run=None -benchmem -timeout 1h ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkPostgresInsert-8            100        1917738232 ns/op        41928271 B/op     693622 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      191.895s
go test -bench=BenchmarkPostgresQuery -benchtime=100x -run=None -benchmem  -timeout 1h  ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkPostgresQuery-8             100           6936319 ns/op           11431 B/op          5 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      0.924s
```

### bench-clickhouse-1000
```
$ make bench-clickhouse-1000
2021/09/09 10:00:43 OK    001_basics.sql
2021/09/09 10:00:43 goose: no migrations to run. current version: 1
go test -bench=BenchmarkClickHouseInsert -benchtime=1000x -run=None -benchmem -timeout 2h ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkClickHouseInsert-8         1000         141200907 ns/op        35162619 B/op     738508 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      141.708s
go test -bench=BenchmarkClickHouseQuery -benchtime=1000x -run=None -benchmem -timeout 2h ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkClickHouseQuery-8          1000          28456499 ns/op            8345 B/op         75 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      28.779s
```

### bench-postgres-1000

```
$ make bench-postgres-1000
2021/09/09 10:46:39 OK    001_basics.sql
2021/09/09 10:46:39 goose: no migrations to run. current version: 1
go test -bench=BenchmarkPostgresInsert -benchtime=1000x -run=None -benchmem -timeout 2h ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkPostgresInsert-8           1000        1864550944 ns/op        41084100 B/op     693606 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      1864.889s
go test -bench=BenchmarkPostgresQuery -benchtime=1000x -run=None -benchmem  -timeout 2h  ./driver/...
goos: darwin
goarch: arm64
pkg: github.com/josudoey/bench-pg-ch/driver
BenchmarkPostgresQuery-8            1000          25294512 ns/op            1250 B/op          4 allocs/op
PASS
ok      github.com/josudoey/bench-pg-ch/driver      25.785s
```

## Disk Usage
### Clickhouse

```
$ clickhouse-client
ClickHouse client version 21.10.1.8000 (official build).
Connecting to localhost:9000 as user default.
Connected to ClickHouse server version 21.10.1 revision 54449.

josudoey.local :) SELECT database, table, count() as number_of_parts, sum(bytes) as sum_size FROM system.parts WHERE active GROUP BY database,table;
SELECT
    database,
    table,
    count() AS number_of_parts,
    sum(bytes) AS sum_size
FROM system.parts
WHERE active
GROUP BY
    database,
    table

Query id: a2866c19-74b6-4bf8-98f5-2d8fa52fb85a

┌─database─┬─table─────────────┬─number_of_parts─┬──sum_size─┐
│ default  │ post_measurements │              71 │ 457168313 │
│ default  │ goose_db_version  │               2 │       552 │
└──────────┴───────────────────┴─────────────────┴───────────┘

2 rows in set. Elapsed: 0.005 sec.
josudoey.local :) SELECT count(*) FROM post_measurements;

SELECT count(*)
FROM post_measurements

Query id: 536e3952-afb6-4082-bce0-28d9bfd2434f

┌──count()─┐
│ 36500000 │
└──────────┘

1 rows in set. Elapsed: 0.002 sec.
```

### Postgres

```sql
$ psql --user postgres
psql (13.4)
Type "help" for help.

postgres=# \timing
Timing is on.
postgres=# SELECT pg_total_relation_size('post_measurements');
 pg_total_relation_size
------------------------
             4589027328
(1 row)

Time: 6.750 ms
postgres=# SELECT count(*) FROM post_measurements;
  count
----------
 36500000
(1 row)

Time: 17027.171 ms (00:17.027)
```