package piscine

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return l.Tail
	}
	return l.Tail.Data
}
