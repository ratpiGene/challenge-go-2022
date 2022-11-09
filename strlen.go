package piscine

func StrLen(s string) int {
	count := 0
	for _, x := range s {
		count = count + 1
		x++
	}
	return count
}
