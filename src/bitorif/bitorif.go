package bitorif

import (
	"os"
)

func isNulNameBit(name string) bool {
	if name[0]|0x20 != 'n' {
		return false
	}
	if name[1]|0x20 != 'u' {
		return false
	}
	if name[2]|0x20 != 'l' {
		return false
	}
	return true
}

func isNulNameIf(name string) bool {
	if name[0] != 'n' && name[0] != 'N' {
		return false
	}
	if name[1] != 'u' && name[0] != 'U' {
		return false
	}
	if name[2] != 'l' && name[0] != 'L' {
		return false
	}
	return true
}

// Because os.DevNull is always ASCII, the bit for uppercase can be omitted
// This version would work on any OS
func isNulNameLoop(name string) bool {
	for i, c := range os.DevNull {
		if int32(name[i]|0x20) != c|0x20 {
			return false
		}
	}
	return true
}
