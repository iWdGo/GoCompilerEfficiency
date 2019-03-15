# Switch useless check

Checking switch performance:
- useless check with nop
- skipping check and relying on default
- returning default in the check

Whatever the choice, the switch has no noticeable variation of performance.

`src>go test -bench=. ./switch` 

**Results**

```
go version go1.12 windows/amd64

BenchmarkCheckUselessCase-4     2000000000               0.43 ns/op
BenchmarkReturnDefault-4        2000000000               0.53 ns/op
BenchmarkNoUselessCase-4        2000000000               0.43 ns/op
PASS
```
 
 