package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		if a[i+1] < a[i] {
			a[i+1], a[i] = a[i], a[i+1]
			i = -1
		}
	}
}
