package arraysstruct

import "strings"

// Required for append
type ase struct {
	name  string
	value int
	prop  string
}
type ame struct {
	value int
	prop  string
}

// Purpose is to set all the words in the order of the int
var (
	names = []string{"a", "b", "c", "d", "e", "f", "g", "h"} // no duplicate
	value = []int{2, 6, 7, 1, 5, 3, 8, 4}
	prop  = []string{"is", "fill", "the", "this", "to", "only", "array", "data"}

	as = []ase{}

	am = map[string]struct {
		value int
		prop  string
	}{}
)

func fill() {
	for i, n := range names {
		as = append(as, ase{n, value[i], prop[i]})
	}

	for i, n := range names {
		am[n] = ame{value[i], prop[i]}
	}
}

func Arrays() string {
	phrase := make([]string, len(names))
	for i, _ := range names {
		phrase[value[i]-1] = prop[i]
	}
	return strings.Join(phrase, " ")
}
func ArrayStructs() string {
	phrase := make([]string, len(as))
	for _, m := range as {
		phrase[m.value-1] = m.prop
	}
	return strings.Join(phrase, " ")
}
func MapStructs() string {
	phrase := make([]string, len(am))
	for _, m := range am {
		phrase[m.value-1] = m.prop
	}
	return strings.Join(phrase, " ")
}
