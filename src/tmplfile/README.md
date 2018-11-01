# Parsing form

Loading template from a file is costly.
It is worth holding it in a global var only for this purpose.

`src>go test -bench=. ./tmplfiles` 

**Results**
```
goos: windows
goarch: amd64
BenchmarkFormParse-4               50000             23246 ns/op
BenchmarkFormParseGlobal-4        500000              3384 ns/op
PASS
```
 