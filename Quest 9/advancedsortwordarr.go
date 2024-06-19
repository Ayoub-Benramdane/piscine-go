package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if Compare(a[j], a[i]) == -1 {
				a[j], a[i] = a[i], a[j]
			}
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
