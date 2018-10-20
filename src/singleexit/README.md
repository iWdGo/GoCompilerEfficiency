# Single exit

In a helper, you might consider to have a single exit point for clarity purpose.

It creates over complex code easily replaced by `range`.
Your cleaner code runs at the same speed when the sole variable is the parameter.
A little disadvantage when found by the end of the range.

`src>go test -bench=. ./singleexit` 

**Results**

```
goos: windows
goarch: amd64
BenchmarkFindStringShort-4              100000000               15.4 ns/op
BenchmarkFindStringRangeShort-4         100000000               15.9 ns/op
BenchmarkFindStringLong-4               20000000                73.1 ns/op
BenchmarkFindStringRangeLong-4          20000000                89.7 ns/op
BenchmarkFindStringNotFound-4           20000000                97.3 ns/op
BenchmarkFindStringRangeNotFound-4      20000000                96.0 ns/op
PASS
```
 
 