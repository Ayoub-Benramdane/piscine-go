package piscine

func StringToIntSlice(str string) []int {
	var res []int
	for _, char := range str {
		res = append(res, int(char))
	}
	return res
}
