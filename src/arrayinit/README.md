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
go version go1.11.2 windows/amd64

BenchmarkInitArrayStatic-4              2000000000               0.40 ns/op
BenchmarkInitArrayMake-4                20000000                57.2 ns/op
BenchmarkInitArrayAppendStatic-4         5000000               273 ns/op
BenchmarkInitArrayAppend-4               5000000               282 ns/op
PASS
```
 
 