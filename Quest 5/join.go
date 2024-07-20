package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i := 0; i < len(strs); i++ {
		res += strs[i]
		if i < len(strs)-1 {
			res += sep
		}
	}
	return res
}
