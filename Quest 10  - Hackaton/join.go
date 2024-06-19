package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i, c := range strs {
		if i == len(strs)-1 {
			res += c
		} else {
			res += c + sep
		}
	}
	return res
}
