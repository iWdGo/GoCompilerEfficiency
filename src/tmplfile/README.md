# Parsing form

Parsing a form using .Parseform requires to build a template.
This template which has no internal logic has to be build on every request.

It is worth holding it in a global var only for this purpose.

`src>go test -bench=. ./templates` 

**Results**
```
goos: windows
goarch: amd64
BenchmarkFormParse-4               50000             23246 ns/op
BenchmarkFormParseGlobal-4        500000              3384 ns/op
PASS
```
 