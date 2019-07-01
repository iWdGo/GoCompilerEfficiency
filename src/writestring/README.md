# Elementary writing of a string to a buffer

The target is of type *buffer.Bytes and the string is simply written to it.
An elementary case is considered. Any formatting or processing makes this benchmark irrelevant.
As `fmt` offers more possibilities, it takes about 3x the cost as io.

Because you need to re-create the buffer from the string using the bytes package, there is no benefit to use the
bytes package.


`src>go test -bench=. ./writestring` 

**Results**

```
go version go1.12.6 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/writestring
BenchmarkFmtWriteString-4        1000000              1962 ns/op
BenchmarkIoWriteString-4         2000000               796 ns/op
BenchmarkBufferWriteTo-4         1000000              1751 ns/op
PASS
```
 
 