package piscine

func BasicAtoi2(s string) int {
	result := 0
	copy := []byte(s)
	count := 1
	for _, r := range s {
		if testNumber(rune(r)) == 0 {
			return result
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		result = result + (int(copy[i])-48)*count
		count = count * 10
	}
	return result
}
