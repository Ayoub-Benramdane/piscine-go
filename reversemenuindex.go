package piscine

func ReverseMenuIndex(menu []string) []string {
	var str []string
	for i := len(menu) - 1; i >= 0; i-- {
		str = append(str, menu[i])
	}
	return str
}
