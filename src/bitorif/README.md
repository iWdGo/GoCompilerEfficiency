# Use bit masking instead of a second if is a marginal improvement

Bit masking is more compact and marginally more efficient than a second if although a bit less readable.

`src>go test -bench=. ./bitorif`

**Results**

```
go version go1.13beta1 windows/amd64

pkg: github.com/iWdGo/GoCompilerEfficiency/src/bitorif
BenchmarkBitMasking-4           1000000000               0.598 ns/op
BenchmarkSecondIf-4             1000000000               0.693 ns/op
PASS
```

 
 
 