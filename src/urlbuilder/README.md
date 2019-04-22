# Url builder

Building a URL is best done using url.URL type. If the operation is done repeatedly,
returning the string or the url is equivalent and the *url.URL should be prefered.

`src>go test -bench=. ./urlbuilder` 

**Results**

```
go version go1.12.2 windows/amd64
pkg: github.com/iWdGo/GoCompilerEfficiency/src/urlbuilder
BenchmarkGetAppString-4          1000000              1227 ns/op
BenchmarkGetAppUrl-4             1000000              1221 ns/op
PASS
```
 
 