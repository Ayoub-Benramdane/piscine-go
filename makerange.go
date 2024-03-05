package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	result := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		result[i] = min + i
	}

	return result
}
