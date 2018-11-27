package isvalid

// Criteria to pass, i.e. 1 to 9 is true
func isValidValue(a int) bool {
	return a < 10
}
func isPositive(a int) bool {
	return a > 0
}

func CheckValueFuncVar(test []int) (s []int) {
	b := false
	if b = isValidValue(test[0]); b && isPositive(test[0]) {
		s = append(s, test[0])
	}
	if b = isValidValue(test[1]); b && isPositive(test[1]) {
		s = append(s, test[1])
	}
	if b = isValidValue(test[2]); b && isPositive(test[2]) {
		s = append(s, test[2])
	}
	if b = isValidValue(test[3]); b && isPositive(test[3]) {
		s = append(s, test[3])
	}
	if b = isValidValue(test[4]); b && isPositive(test[4]) {
		s = append(s, test[4])
	}
	return
}

func CheckValueIfVar(test []int) (s []int) {
	if b := isValidValue(test[0]); b && isPositive(test[0]) {
		s = append(s, test[0])
	}
	if b := isValidValue(test[1]); b && isPositive(test[1]) {
		s = append(s, test[1])
	}
	if b := isValidValue(test[2]); b && isPositive(test[2]) {
		s = append(s, test[2])
	}
	if b := isValidValue(test[3]); b && isPositive(test[3]) {
		s = append(s, test[3])
	}
	if b := isValidValue(test[4]); b && isPositive(test[4]) {
		s = append(s, test[4])
	}
	return
}
func CheckValueForFunc(test []int) (s []int) {
	b := false
	for _, v := range test {
		if b = isValidValue(v); b && isPositive(v) {
			s = append(s, v)
		}
	}
	return
}
func CheckValueForIf(test []int) (s []int) {
	for _, v := range test {
		if b := isValidValue(v); b && isPositive(v) {
			s = append(s, v)
		}
	}
	return
}
