package piscine

func Map(f func(int) bool, a []int) []bool {
	table := make([]bool, len(a))
	for i, c := range a {
		table[i] = f(c)
	}
	return table
}
