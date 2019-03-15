# Data in multiple arrays

You need to maintain related data using an ordinary string key:
- Each array contains one data of one set
- Array of struct
- Map using the key to access the struct

`src>go test -bench=. ./arraysstruct` 

**Results**

Arrays are slightly faster.
Performance of an array of struct is similar with the advantage of stronger typing.
The map is 30% more expensive.

```
go version go12 windows/amd64

BenchmarkArrays-4                5000000               285 ns/op
BenchmarkArrayStructs-4          5000000               300 ns/op
BenchmarkMapStructs-4            3000000               413 ns/op
PASS
```
