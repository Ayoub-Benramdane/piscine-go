package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		if i < len(a)-1 && Compare(a[i+1], a[i]) == -1 {
			a[i+1], a[i] = a[i], a[i+1]
			i = -1
		}
	}
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}
