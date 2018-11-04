package arrayinit

// Luckily, it is only 10 items...
func InitArrayStatic() []int {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return a
}
func InitArrayMake() []int {
	a := make([]int, 10)
	for i := 1; i <= 10; i++ {
		a[i-1] = i
	}
	return a
}
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
