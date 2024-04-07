package piscine

func Swap(a *int, b *int) {
	x := *a
	*a = *b
	*b = x
}
