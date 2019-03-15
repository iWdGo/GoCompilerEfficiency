# Don't keep length of array elsewhere

Len() for an array is as efficient as any other method.

`src>go test -bench=. ./lenarray`

**Results**

```
go version go1.12 windows/amd64

BenchmarkLenOfArray-4             200000              6096 ns/op
BenchmarkLenOfArrayVar-4          300000              5916 ns/op
PASS   
```

 
 
 