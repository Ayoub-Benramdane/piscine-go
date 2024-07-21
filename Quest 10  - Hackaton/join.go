package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i, c := range strs {
		res += c
		if i < len(strs)-1 {
			res += sep
		}
	}
	return res
}
