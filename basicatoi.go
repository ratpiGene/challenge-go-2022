package piscine

func testNumber(try rune) int {
	for i := '0'; i <= '9'; i++ {
		if try == i {
			return 1
		}
	}
	return 0
}

func BasicAtoi(s string) int {
	result := 0
	count := 1
	safe := []byte(s)
	for _, r := range s {
		if testNumber(rune(r)) == 0 {
			return result
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		result = result + (int(safe[i])-48)*count
		count = count * 10
	}
	return result
}
