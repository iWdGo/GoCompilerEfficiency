package emptycase

func CheckUselessCase(a int) int {
	switch a {
	case 0:
		return 10
	case 1:
		return 8
	case 2:
		// Skipping
	case 3:
		return 2
	}
	return -1
}

func CheckReturnDefault(a int) int {
	switch a {
	case 0:
		return 10
	case 1:
		return 8
	case 2:
		return -1
	case 3:
		return 2
	}
	return -1
}
func CheckNoUselessCase(a int) int {
	switch a {
	case 0:
		return 10
	case 1:
		return 8
	case 3:
		return 2
	}
	return -1
}
