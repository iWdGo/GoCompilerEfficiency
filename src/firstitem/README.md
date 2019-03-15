# First item in a loop

When you build a result using a for loop, the first item might need a specific treatment.

The most efficient solution is to use an `if` statement to avoid code duplication.
Whatever the test, no significant difference appears.

The unfair solution doesn't take the full hypothesis into account and gets a 10% improvement.

`src>go test -bench=. ./firstitem` 

**Results**

```
go version go1.12 windows/amd64

BenchmarkFirstItemBefore-4                 20000             86551 ns/op
BenchmarkFirstItemIfOnIndex-4              20000             88000 ns/op
BenchmarkFirstItemIfOnArray-4              20000             88107 ns/op
BenchmarkFirstItemBeforeUnfair-4           20000             70460 ns/op
PASS
```
 
 