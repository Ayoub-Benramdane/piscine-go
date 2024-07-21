package piscine

func Compact(ptr *[]string) int {
	table := *ptr
	compact := make([]string, 0)
	for _, val := range table {
		if val != "" {
			compact = append(compact, val)
		}
	}
	*ptr = compact
	return len(compact)
}
