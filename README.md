flatbuffer vs json:
===================

Upsides of flatbuffer:
=====================

* Data integrity by using static typing
* Faster than interpreted formats i.e. json
* True random access like json

Sample run:
==========

```
-> % flatc -g schema/dcpevent.fbs
-> % go test -bench=’.*’ ./
testing: warning: no tests to run
BenchmarkFlatbufRead-8          20000000                74.7 ns/op       964.05 MB/s           0 B/op          0 allocs/op
BenchmarkFlatbufWrite-8          5000000               288 ns/op         249.31 MB/s           0 B/op          0 allocs/op
BenchmarkJSONRead-8              2000000               759 ns/op          18.43 MB/s         256 B/op          3 allocs/op
BenchmarkJSONWrite-8             3000000               517 ns/op          42.52 MB/s         264 B/op          5 allocs/op
PASS
ok      github.com/abhi-bit/flatbuf-vs-json     7.641s
```
