package piscine

func Abort(a, b, c, d, e int) int {
	x := []int{a, b, c, d, e}
	for i := 0; i < len(x); i++ {
		if i < len(x)-1 && x[i] > x[i+1] {
			x[i], x[i+1] = x[i+1], x[i]
			i = -1
		}
	}
	return x[2]
}
