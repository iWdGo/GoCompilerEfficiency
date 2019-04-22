// Arrayinit initializes an array in various ways including the unfair static method
package arrayinit

// Array is known
func InitArrayStatic() []int {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return a
}

// Array is made and filled
func InitArrayMake() []int {
	a := make([]int, 10)
	for i := 1; i <= 10; i++ {
		a[i-1] = i
	}
	return a
}

// Var is created and values appended
func InitArrayAppend() []int {
	a := make([]int, 0)
	for i := 1; i <= 10; i++ {
		a = append(a, i)
	}
	return a
}

// Appending on a static var
func InitArrayAppendStatic() []int {
	var a []int
	for i := 1; i <= 10; i++ {
		a = append(a, i)
	}
	return a
}
