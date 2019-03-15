# Parsing form

Parsing a form using .Parseform requires to build a template.
This template which has no internal logic has to be build on every request.

It is worth holding it in a global var only for this purpose.

`src>go test -bench=. ./formparse` 

**Results**
```
go version go1.12 windows/amd64

BenchmarkFormParse-4              100000             22177 ns/op
BenchmarkFormParseGlobal-4        500000              3054 ns/op
PASS
```
 