package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	items := SplitWhiteSpaces(str)
	for _, i := range items {
		summary[i]++
	}
	return summary
}

func SplitWhiteSpaces(str string) []string {
	res := []string{}
	s := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			s += string(str[i])
		} else if s != "" {
			res = append(res, s)
			s = ""
		}
	}
	if s != "" {
		res = append(res, s)
	}
	return res
}
