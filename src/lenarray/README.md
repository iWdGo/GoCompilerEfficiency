# Don't keep length of array elsewhere

Len() for an array is as efficient as any other method.

Keeping a separate variable in math functions for instance has no use.

`src>go test -bench=. ./lenarray`

**Results**

```
goos: windows
goarch: amd64
BenchmarkLenOfArray-4             200000              6147 ns/op
BenchmarkLenOfArrayVar-4          200000              6316 ns/op
PASS   
```

 
 
 