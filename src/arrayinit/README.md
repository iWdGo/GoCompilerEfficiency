# Initializing an array

Initializing an array can be:
- static
- with a known length and make
- using append

Comparison is unfair but sheds some light on the cost.
The most efficient is obviously static. 
Append is extending the array and x5 the make().

The method may depend on your app requirement and/or practical constraints.


`src>go test -bench=. ./arrayinit` 

**Results**

```
go version go1.12 windows/amd64

BenchmarkInitArrayStatic-4              2000000000               0.49 ns/op
BenchmarkInitArrayMake-4                20000000                54.6 ns/op
BenchmarkInitArrayAppendStatic-4         5000000               266 ns/op
BenchmarkInitArrayAppend-4               5000000               277 ns/op
PASS
```
 
 