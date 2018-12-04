# Switch useless check

Checking switch performance:
- useless check with nop
- skipping check and relying on default
- returning default in the check

Whatever the choice, the switch has no noticeable variation of performance.

`src>go test -bench=. ./switch` 

**Results**

```
go version go1.11.2 windows/amd64
BenchmarkCheckUselessCase-4     2000000000               0.47 ns/op
BenchmarkReturnDefault-4        2000000000               0.47 ns/op
BenchmarkNoUselessCase-4        2000000000               0.49 ns/op
PASS
```
 
 