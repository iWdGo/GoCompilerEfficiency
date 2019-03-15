# Appending elements to an array of array

Adding a element to an array is done using append.
The array is defined as:
- [][]int which is always the underlying type
- []t
- tt where a method is used

Using a method is the most efficient even if difference can be small.

`src>go test -bench=. ./arrayofarray`

**Results**

```
go version go1.12 windows/amd64
```
add one cell (inc == 1)
```
BenchmarkUpdateArray-4          10000000               165 ns/op
BenchmarkUpdateArrayType-4      20000000               130 ns/op
BenchmarkUpdateArrayMethod-4    20000000               113 ns/op
```

add 10 cells (inc == 10)
```
BenchmarkUpdateArray-4           1000000              1751 ns/op
BenchmarkUpdateArrayType-4       1000000              1297 ns/op
BenchmarkUpdateArrayMethod-4     1000000              1008 ns/op
PASS
```
 
 