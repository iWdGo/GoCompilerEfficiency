# Get bit representation of a byte

When based on strings, `%b` directive of the `fmt` package provides a bit display.
But it requires to range on chars which is prohibitive compared to other methods.

Rotating the bits each time introduces some penalty.

The simplest code is using the `bits` package and is marginally more efficient.

`src>go test -bench=. ./bytetobits`

**Results**

```
go version go1.13 windows/amd64

pkg: github.com/iWdGo/GoCompilerEfficiency/src/bytetobits
BenchmarkRangeToBits-4           2522445               472 ns/op
BenchmarkBits-4                 11327166               101 ns/op
BenchmarkAndToBits-4            13191872                93.4 ns/op
BenchmarkBitsToBits-4           12252073                90.8 ns/op
PASS
```

 
 
 