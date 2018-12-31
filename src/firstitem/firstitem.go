package firstitem

import (
	"strconv"
)

func FirstItemBefore(a int) (l int) {
	res := []string{"1"}
	for i := 2; i < 1000; i++ {
		res = append(res, "+")
		res = append(res, strconv.Itoa(i))
	}
	return len(res)
}

func FirstItemIfOnIndex(a int) (l int) {
	res := make([]string, 0)
	for i := 1; i < a; i++ {
		if i != 1 { // specific treatment
			res = append(res, "+")
		}
		res = append(res, strconv.Itoa(i))
	}
	return len(res)
}

func FirstItemIfOnArray(a int) (l int) {
	res := make([]string, 0)
	for i := 1; i < a; i++ {
		if len(res) != 0 { // specific treatment
			res = append(res, "+")
		}
		res = append(res, strconv.Itoa(i))
	}
	return len(res)
}

func FirstItemBeforeUnfair(a int) (l int) {
	res := []string{"1"}
	for i := 2; i < a; i++ {
		// This is not representative as the hypothesis implies that the first
		// item has some complex treatment creating code duplication.
		res = append(res, "+", strconv.Itoa(i))
	}
	return len(res)
}
