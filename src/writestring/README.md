# Elementary writing of a string to a buffer

The target is of type *buffer.Bytes and the `string` is simply written to it.
An elementary case is considered. Any formatting or processing makes this benchmark irrelevant.
This module is using methods writing a `string` without conversion.
 - `fmt.Fprint`
 - `WriteString()` method of `buffer` package
 - `io.WriteString()`   

As `fmt` offers more possibilities, it takes about 3x the cost as others.

`src>go test -bench=. ./writestring` 

**Results**

```
go version go1.12.6 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/writestring
BenchmarkFmtWriteString-4        1000000              1885 ns/op
BenchmarkIoWriteString-4         2000000               694 ns/op
BenchmarkBufferWriteString-4     3000000               447 ns/op
PASS
```
 
 