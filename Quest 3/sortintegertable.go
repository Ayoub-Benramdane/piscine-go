package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[i] {
				table[j], table[i] = table[i], table[j]
			}
		}
	}
}
