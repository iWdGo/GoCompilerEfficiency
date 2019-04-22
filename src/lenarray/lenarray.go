package lenarray

import "fmt"

// Global var
func lenOfArray() (res int) {
	a := make([]int, 1)
	for i := 1; i < 1000; i++ {
		a = append(a, len(a))
		if len(a) != i+1 {
			fmt.Println("check you code")
		}
	}
	return a[len(a)-1] // compiler cannot optimize the full loop
}

// Local var
func lenOfArrayVar() (res int) {
	a := make([]int, 1)
	la := 1
	for i := 1; i < 1000; i++ {
		a = append(a, la)
		if la != i {
			fmt.Println("check you code")
		}
		la++
	}
	return a[la-1] // compiler cannot optimize the full loop
}
