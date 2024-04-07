package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index <= 1 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
