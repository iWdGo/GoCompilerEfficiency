package bitorif

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
