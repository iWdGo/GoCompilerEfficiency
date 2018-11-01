# Content to File

Most efficient structure to write a file from some content.
No difference between a string and *bytes.Buffer.

`src>go test -bench=. ./tofile` 

**Results**
```
goos: windows
goarch: amd64
BenchmarkStringToFile-4              300           4887857 ns/op
BenchmarkBufferToFile-4              300           4717306 ns/op
PASS
```
 