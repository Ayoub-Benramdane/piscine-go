package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
