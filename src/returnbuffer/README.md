# Returning bytes.Buffer content

Content can be returned as:
- a buffer
- as an array of bytes
- as a string

Returning a string from the buffer is marginally more expensive as expected.

`src>go test -bench=. ./returnbuffer` 

**Results**

```
go version go1.12.6 windows/amd64

pkg: github.com/iWdGo/GoCompilerEfficiency/src/returnbuffer
BenchmarkReturnBuffer-4                   300000              4367 ns/op
BenchmarkReturnBufferString-4             300000              4570 ns/op
BenchmarkReturnBufferBytes-4              300000              4360 ns/op
PASS
```
 
 