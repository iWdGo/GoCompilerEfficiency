package main

import (
	"fmt"
	"math/bits"
)

// fmt package and %b directive.
// strings.SplitN has no use as []string is returned.
func rangeToBits(data []byte) (st []int) {
	st = make([]int, len(data)*8) // Performance x 2 as no append occurs
	for i, d := range data {
		// Sprintf is overkill as you get 48 (0) and 49 (1)
		for j, c := range fmt.Sprintf("%08b", d) {
			if c == 48 { // "0"
				st[i*8+j] = 0
			} else {
				st[i*8+j] = 1
			}
		}
	}
	return
}

// Isolating each bit in one rotation.
func forToBits(bs []byte) []int {
	r := make([]int, len(bs)*8)
	for i, b := range bs {
		for j := 0; j < 8; j++ {
			r[i*8+j] = int(b >> uint(7-j) & 0x01)
		}
	}
	return r
}

// Using and operation to isolate the bit and no rotation.
func andToBits(data []byte) (st []int) {
	st = make([]int, len(data)*8) // Performance x 2 as no append occurs.
	var j uint
	exp := 0
	for i, d := range data {
		exp = 0
		j = 256 // 2 with exponent 8 to get leftmost bit first.
		for exp < 8 {
			j /= 2
			if b := uint(d) & j; b == j {
				st[i*8+exp] = 1
			} else {
				st[i*8+exp] = 0
			}
			exp += 1
		}
	}
	return
}

// Using the bits package and a rotation.
func bitsToBits(data []byte) (st []int) {
	st = make([]int, len(data)*8) // Performance x 2 as no append occurs.
	for i, d := range data {
		for j := 0; j < 8; j++ {
			if bits.LeadingZeros8(d) == 0 {
				// No leading 0 means that it is a 1
				st[i*8+j] = 1
			} else {
				st[i*8+j] = 0
			}
			d = d << 1
		}
	}
	return
}
