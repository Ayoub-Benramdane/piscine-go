package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	items := SplitWhiteSpace(str)

	for _, i := range items {
		summary[i]++
	}

	return summary
}

func SplitWhiteSpace(s string) []string {
	x := []rune(s)
	var y []string
	a := ""

	for i := 0; i < len(x); i++ {
		if x[i] != ' ' {
			a += string(x[i])
		} else {
			if i < len(x) {
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
