# Reading each byte of a buffer

Each byte (`bg`) can be read using:
- `_, _ = bb.Read(bg)`
- `bg, _ = bb.ReadByte()`

When treating every byte, `ReadByte` provides 20%+ improvement.
One less variable is part of the explanation.

`src>go test -bench=. ./returnbyte` 

**Results**

```
go version go1.12.6 windows/amd64

pkg: github.com/iWdGo/GoCompilerEfficiency/src/readbyte
BenchmarkBufferRead-4             200000              8540 ns/op
BenchmarkRBufferReadByte-4        200000              5971 ns/op
PASS
```
 
 