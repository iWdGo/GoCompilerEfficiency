# Content to File

Most efficient structure to write a file from some content.
No difference between a string and *bytes.Buffer.

`src>go test -bench=. ./tofile` 

**Results**
```
go version go1.12 windows/amd64

BenchmarkStringToFile-4             3000            411431 ns/op
BenchmarkBufferToFile-4             3000            423128 ns/op
PASS
```
 