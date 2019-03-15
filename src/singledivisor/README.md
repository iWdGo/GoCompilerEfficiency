# Is 2 the sole divisor ?

The most efficient way to check that 2 is the sole divisor is to use math.Log2()

Algorithms:
- math.Log2() is the most efficient and precise enough for int
- Shifting the int is a close second.
- Using standard division between the int and 2 is already 10% behind.
- Usual log division is not precise enough for larger int's but is only 2x the expense.
- Using .IsInt() and the related conversion from math/big is costly.
- Standard division is the most generic must is also the most expensive.

`src>go test -bench=. ./singledivisor` 

**Results**

```
go version go1.12 windows/amd64

BenchmarkIs2SingleDivLog-4              100000000               10.8 ns/op
BenchmarkIs2SingleDivShift-4            100000000               13.4 ns/op
BenchmarkIs2SingleDivDivision-4         100000000               15.0 ns/op
BenchmarkIsSingleDivLog-4               30000000                36.7 ns/op
BenchmarkIs2SingleDivBig-4              20000000                70.8 ns/op
BenchmarkIsSingleDivDivision-4          10000000               173 ns/op
PASS   
```

 
 
 