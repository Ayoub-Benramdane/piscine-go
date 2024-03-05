package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			res = res + strs[i] + sep
		} else {
			res = res + strs[i]
		}
	}
	return res
}
