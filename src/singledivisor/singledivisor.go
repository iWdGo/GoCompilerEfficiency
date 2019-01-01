package forloops

import (
	"math"
	"math/big"
)

// 2 is the divisor
func is2SingleDivBig(a int) (b bool) {
	return big.NewFloat(math.Log2(float64(a))).IsInt()
}
func is2SingleDivLog(a int) (b bool) {
	l := math.Log2(float64(a))
	return (l - math.Floor(l)) == 0 // No decimal part
}
func is2SingleDivDivision(a int) (b bool) {
	for a%2 == 0 {
		a = a / 2
	}
	return a == 1 // 1 means division is complete
}
func is2SingleDivShift(a int) (b bool) {
	for a%2 == 0 {
		a = a >> 1
	}
	return a == 1 // 1 means division is complete
}

// Any divisor
func isSingleDivLog(a, d int) (b bool) {
	l := (math.Log(float64(a)) / math.Log(float64(d)))
	return (l - math.Floor(l)) == 0 // No decimal part
}
func isSingleDivDivision(a, d int) (b bool) {
	for a%d == 0 {
		a = a / d
	}
	return a == 1
}
