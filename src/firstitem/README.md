# First item in a loop

When you build a result using a for loop, the first item might need a specific treatment.

The most efficient solution is to use an `if` statement to avoid code duplication.
Whatever the test, no significant difference appears.

The unfair solution doesn't take the full hypothesis into account and gets a 10% improvement.

`src>go test -bench=. ./firstitem` 

**Results**

```
go version go1.11.4 windows/amd64

BenchmarkFirstItemBefore-4                 20000             87899 ns/op
BenchmarkFirstItemIfOnIndex-4              20000             88300 ns/op
BenchmarkFirstItemIfOnArray-4              20000             88859 ns/op
BenchmarkFirstItemBeforeUnfair-4           20000             72358 ns/op
PASS
```
 
 