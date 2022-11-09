package piscine

func AppendRange(min, max int) []int {
	size := max - min
	a := min
	var answer []int
	if size > 0 {
		for i := 0; i < size; i++ {
			answer = append(answer, a)
			a = a + 1
		}
		return answer
	} else {
		for i := 0; i < 1; i++ {
			i = i + 1
		}
	}
	return answer
}
