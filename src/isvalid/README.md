# Local boolean

Checking a sequence of boolean conditions:
- individually with a boolean in func scope
- individually with a boolean in if scope
- inside a for loop using a bool in func scope
- inside a for loop shadowing a bool in if scope

The example is theoretical and should prevent compiler optimization of the tests applied to each value.

If the `if` instructions are in sequence, var creation inside the if is expensive.
Inside a for loop for a sequence of tests of the same size, shadowing is so optimized that
you don't need to bother.


`src>go test -bench=. ./isvalid` 

**Results**

```
go version go1.11.2 windows/amd64
BenchmarkCheckValueFuncVar-4    20000000                82.5 ns/op
BenchmarkCheckValueIfVar-4      20000000               113 ns/op
BenchmarkCheckValueForFunc-4    10000000               107 ns/op
BenchmarkCheckValueForIf-4      20000000                97.7 ns/op
PASS
```
 
 