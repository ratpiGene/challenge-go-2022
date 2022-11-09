package piscine

func ConcatParams(args []string) string {
	b := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			b = b + args[i]
		} else if i < len(args)-1 {
			b = b + args[i] + "\n"
		}
	}
	return b
}
