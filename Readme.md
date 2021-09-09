# benchmark postgres clickhouse

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