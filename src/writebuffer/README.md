# Elementary write of a buffer to a buffer

Writing a buffer to another buffer can be donc using
 - `fmt.Fprint`
 - `WriteTo()` method

As `fmt` offers more possibilities, it is about 6x heavier than io.

`>go test -bench=. ./src/writebuffer` 

**Results**

```
go version go1.12.6 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/writebuffer
BenchmarkFmtWriteString-4         300000              3801 ns/op
BenchmarkIoWriteString-4         2000000               688 ns/op
PASS
```
 