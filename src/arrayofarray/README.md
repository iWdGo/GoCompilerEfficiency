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
go version go1.12beta1 windows/amd64
```
add one cell (inc == 1)
```
BenchmarkUpdateArray-4          10000000               231 ns/op
BenchmarkUpdateArrayType-4      20000000               154 ns/op
BenchmarkUpdateArrayMethod-4    20000000               119 ns/op
```

add 10 cells (inc == 10)
```
BenchmarkUpdateArray-4           1000000              1872 ns/op
BenchmarkUpdateArrayType-4       1000000              1220 ns/op
BenchmarkUpdateArrayMethod-4     1000000              1116 ns/op
PASS
```
 
 