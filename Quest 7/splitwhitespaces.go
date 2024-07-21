package piscine

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
