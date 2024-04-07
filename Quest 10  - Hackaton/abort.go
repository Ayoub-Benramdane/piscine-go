package piscine

func Abort(a, b, c, d, e int) int {
	x := []int{a, b, c, d, e}
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	return x[2]
}
