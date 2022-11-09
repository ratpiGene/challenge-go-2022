package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	var b []int
	x := 0
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			b = append(b, i)
		}
		if i == len(s)-1 {
			b = append(b, i+1)
		}
	}
	for i := 0; 1 < len(b); i++ {
		for y := x; y < b[i]; i++ {
			if s[y] != 32 {
				word = word + string(s[y])
			}
		}
		if len(word) > 0 {
			answer = append(answer, word)
			word = ""
			x = b[i] + 1
		}
	}
	return answer
}
