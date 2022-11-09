package piscine

func StrRev(s string) string {
	reverse := []byte(s)
	safe := []byte(s)
	for i := 0; i < len(s); i++ {
		safe[(len(s)-1)-i] = reverse[i]
	}
	s = string(safe)
	return s
}
