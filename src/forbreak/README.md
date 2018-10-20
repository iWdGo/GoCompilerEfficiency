# For loops should use a local var

For loop construction is very optimized in Go and for loops are everywhere.

The most efficient loop is using a locally defined variable.

`src>go test -bench=. ./forbreak` 

**Results**

```
goos: windows
goarch: amd64
BenchmarkForGlobalVar-4           500000              2568 ns/op
BenchmarkForLocalVar-4           2000000               574 ns/op
PASS   
```

 
 
 