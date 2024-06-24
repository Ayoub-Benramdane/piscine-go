package piscine

func ListSize(l *List) int {
	i := 0
	for l.Head != nil {
		i++
		l.Head = l.Head.Next
	}
	return i
}
