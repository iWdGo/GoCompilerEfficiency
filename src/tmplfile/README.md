# Parsing form

Loading template from a file is costly.
It is worth holding it in a global var only for this purpose.

`src>go test -bench=. ./tmplfiles` 

**Results**
```
go version go1.12 windows/amd64

BenchmarkLoadTemplateFile-4                 3000            376455 ns/op
BenchmarkLoadTemplateFileOnce-4           200000              7170 ns/op
PASS
```
 