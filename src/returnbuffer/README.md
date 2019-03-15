# Returning bytes.Buffer content

Content can be returned as:
- a buffer
- as an array of bytes
- as a string

Returning a string from the buffer is marginally more expensive as expected.

`src>go test -bench=. ./returnbuffer` 

**Results**

```
go version go1.12 windows/amd64

BenchmarkReturnBuffer-4                   300000              4390 ns/op
BenchmarkReturnBufferString-4             300000              4584 ns/op
BenchmarkReturnBufferBytes-4              300000              4380 ns/op
PASS
```
 
 