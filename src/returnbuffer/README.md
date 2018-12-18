# Returning bytes.Buffer content

Content can be returned as:
- a buffer
- as an array of bytes
- as a string

Returning a string from the buffer is marginally more expensive as expected.

`src>go test -bench=. ./returnbuffer` 

**Results**

```
go version go1.11.4 windows/amd64

BenchmarkReturnBuffer-4                   300000              4567 ns/op
BenchmarkReturnBufferString-4             300000              4640 ns/op
BenchmarkReturnBufferBytes-4              300000              4504 ns/op
PASS
```
 
 