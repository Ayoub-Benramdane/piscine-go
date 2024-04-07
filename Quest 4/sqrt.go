package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	var result int
	for result = 0; result*result <= nb; result++ {
		if result*result == nb {
			return result
		}
	}

	return 0
}
