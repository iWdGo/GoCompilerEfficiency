# Single exit

In a helper, you might consider to have a single exit point for clarity purpose.
This structure produces overly complex code easily replaced by `range`.
In Go 1.11, the code is on average 5% faster except towards the end of the searched range.
With Go 1.12, codes are equivalent as `range` is regularly catching up the `while` structure.

`src>go test -bench=. ./singleexit` 

`go version go1.11.1 windows/amd64`
![while vs range](whilevsrange111.png?raw=true "While vs Range")

`go version go1.12 windows/amd64`
 
 In the box, before the not found case, you can see the cost of exiting a for loop just before its end.
 
![while vs range](whilevsrange.png?raw=true "While vs Range")
 