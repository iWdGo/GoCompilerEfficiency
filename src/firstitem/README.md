# First item in a loop

When you build a result using a for loop, the first item might need a specific treatment.

The most efficient solution is to use an `if` statement to avoid code duplication.
The unfair solution is optimizing w/o taking the full hypothesis into account and gets a 10% improvement.

`src>go test -bench=. ./firstitem` 

**Results**

```
goos: windows
goarch: amd64
BenchmarkFirstItemBefore-4                 20000             90101 ns/op
BenchmarkFirstItemIf-4                     20000             89894 ns/op
BenchmarkFirstItemBeforeUnfair-4           20000             72109 ns/op
PASS
```
 
 