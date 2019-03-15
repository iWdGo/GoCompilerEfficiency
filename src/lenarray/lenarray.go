package lenarray

import "fmt"

// Effect of global var on for loop
func LenOfArray() (res int) {
	a := make([]int, 1)
	for i := 1; i < 1000; i++ {
		a = append(a, len(a))
		if len(a) != i+1 {
			fmt.Println("check you code")
		}
	}
	return a[len(a)-1] // compiler cannot optimize the full loop
}

func LenOfArrayVar() (res int) {
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
