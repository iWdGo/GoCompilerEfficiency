# Reading each byte of a buffer

Each byte (`bg`) can be read using:
- `_, _ = bb.Read(bg)`
- `bg, _ = bb.ReadByte()`

So, if you're treating everybyte, a 30% benefit to use ReadByte.
One less variable is part of the explanation.

`src>go test -bench=. ./returnbyte` 

**Results**

```
goos: windows
goarch: amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/readbyte
BenchmarkBufferRead-4             200000              8540 ns/op
BenchmarkRBufferReadByte-4        200000              5971 ns/op
PASS
```
 
 