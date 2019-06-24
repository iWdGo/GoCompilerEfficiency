# Elementary print of a string

Printing a string using the fmt package is similar to io.WriteString.
As `fmt` offers more possibilities, it takes about 3x the cost as io.

`src>go test -bench=. ./writestring` 

**Results**

```
go version go1.12.2 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/writestring
BenchmarkFmtWriteString-4         500000              2002 ns/op
BenchmarkIoWriteString-4         2000000               635 ns/op
PASS
```
 
 