# Elementary write of a buffer to a buffer

The target is of type *buffer.Bytes and the `buffer` is simply written to it.
An elementary case is considered. Any formatting or processing makes this benchmark irrelevant.
This module is using methods writing a `buffer` without conversion.
 - `fmt.Fprint`
 - `Write()` method of buffer package with a conversion to `bytes`
 - `WriteTo()` method of buffer package
 - `io.CopyBuffer()`   

As `fmt` offers more possibilities, it is about 6x heavier than the fastest alternate solution.
The byte conversion keeps the buffer to copy intact and other methods are thus heavier because
it must be re-filled.

This benchmark is irrelevant if any formatting occurs during writing.

`>go test -bench=. ./src/writebuffer` 

**Results**

```
go version go1.12.6 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/writebuffer
BenchmarkWriteBuffer-4           3000000               481 ns/op
BenchmarkWriteToBuffer-4         1000000              1745 ns/op
BenchmarkIoCopyBuffer-4          1000000              2304 ns/op
BenchmarkFmtWriteBuffer-4         500000              3783 ns/op
PASS
```
 