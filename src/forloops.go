package forbreak

// Effect of global var on for loop
func ForLocalVar() (res int) {
	for i := 1; i < 1000; i++ {
		res = res + i
	}
	return
}

var j int

func ForGlobalVar() (res int) {
	for j = 1; j < 1000; j++ {
		res = res + j
	}
	return
}
