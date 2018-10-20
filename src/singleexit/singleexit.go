package singleexit

import (
	"strings"
)

var vArray = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}

// One exit point and a default value
func FindString(s string) (b bool) {
	b = false // value when not found
	k := 0
	for k < len(vArray) && !b { //range cannot be used as range does not return bool for &&
		b = strings.HasPrefix(s, vArray[k]) //group code found as prefix. stopping search
		k++
	}
	return
}

func FindStringRange(s string) bool {
	for _, c := range vArray { //range cannot be used as range does not return bool for &&
		if strings.HasPrefix(s, c) {
			return true
		}
	}
	return false
}
