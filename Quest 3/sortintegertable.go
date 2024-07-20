package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		if i < len(table)-1 && table[i+1] < table[i] {
			table[i], table[i+1] = table[i+1], table[i]
			i = -1
		}
	}
}
