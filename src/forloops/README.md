# For loops should use a local var

For loop construction is very optimized in Go and for loops are everywhere.

The most efficient loop is using a locally defined variable.

`src>go test -bench=. ./forbreak` 

**Results**

```
go version go1.12 windows/amd64

BenchmarkForGlobalVar-4           500000              2720 ns/op
BenchmarkForLocalVar-4           3000000               567 ns/op
PASS   
```

 
 
 