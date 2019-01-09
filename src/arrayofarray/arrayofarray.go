package arrayofarray

type (
	t  []int
	tt []t
)

const (
	size = 5
)

var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// [][]int
func initArray() [][]int {
	ab := make([][]int, size)
	for i, _ := range ab {
		ab[i] = a[i:]
	}
	return ab
}
func updateArray(ab [][]int, inc int) [][]int {
	for k, _ := range ab {
		for i := 0; i < inc; i++ {
			// b = append(b, i): no update
			// ab[k] = append(b, i): last update
			ab[k] = append(ab[k], i)
		}
	}
	return ab
}

// []t
func initArrayType() []t {
	ab := make([]t, size)
	for i, _ := range ab {
		ab[i] = a[i:]
	}
	return ab
}
func updateArrayType(ab []t, inc int) []t {
	for k, _ := range ab {
		for i := 0; i < inc; i++ {
			ab[k] = append(ab[k], i)
		}
	}
	return ab
}

// tt
func initArrayTyped() tt {
	ab := make([]t, size)
	for i, _ := range ab {
		ab[i] = a[i:]
	}
	return ab
}
func (ab *tt) updateArrayMethod(inc int) {
	abt := *ab
	for k, _ := range abt {
		for i := 0; i < inc; i++ {
			abt[k] = append(abt[k], i)
		}
	}
	*ab = abt
}
