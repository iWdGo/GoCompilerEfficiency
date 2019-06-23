# Parsing form

Parsing a form using .Parseform requires to build a template.
This template which has no internal logic has to be build on every request.
Three options
- build a global template
- re-build the template for every parse
- use implicit parsing

Implicit parsing has a small advantage but reduce the code size significantly.

`src>go test -bench=. ./formparse` 

**Results**
```
go version go1.12.6 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/formparse
BenchmarkFormParse-4              100000             23966 ns/op
BenchmarkFormParseGlobal-4        500000              3158 ns/op
BenchmarkFormParseImplicit-4      500000              2846 ns/op
PASS
```
 