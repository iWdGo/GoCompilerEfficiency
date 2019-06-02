# Using template to update one value

Using a template when an elementary Fprint() suffices is costly.

`src>go test -bench=. ./tmplhtml` 

**Results**

```
go version go1.12.5 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/tmplhtml
BenchmarkHTMLOutputTmpl-4         500000              3717 ns/op
BenchmarkHTMLOutputBuffer-4      5000000               354 ns/op
PASS
```
 
 