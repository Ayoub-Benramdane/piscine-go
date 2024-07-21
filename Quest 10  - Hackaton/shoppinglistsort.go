package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		if len(slice[i]) > len(slice[i+1]) {
			slice[i], slice[i+1] = slice[i+1], slice[i]
			i = -1
		}
	}
	return slice
}
