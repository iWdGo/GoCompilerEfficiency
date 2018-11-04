# Data in multiple arrays

You need to maintain related data using an ordinary string key:
- Each array contains one data of one set
- Array of struct
- Map using the key to access the struct

`src>go test -bench=. ./arraysstruct` 

**Results**

Arrays are slightly the faster.
An array of struct has a similar performance with the advantage of stronger typing.
The map is 30% more expensive.

```
go version go1.11.2 windows/amd64

BenchmarkArrays-4                5000000               331 ns/op
BenchmarkArrayStructs-4          5000000               345 ns/op
BenchmarkMapStructs-4            3000000               464 ns/op
PASS
```
