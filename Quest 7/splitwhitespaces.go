package piscine

func SplitWhiteSpaces(str string) []string {
	x := []rune(str)
	var y []string
	a := ""

	for i := 0; i < len(x); i++ {
		if x[i] != ' ' {
			a += string(x[i])
		} else {
			if a != "" && a != " " && i < len(x) && x[i] == ' ' {
				y = append(y, a)
			}
			a = ""
		}
	}
	if a != "" {
		y = append(y, a)
	}
	return y
}
