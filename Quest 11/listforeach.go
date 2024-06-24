package piscine

func ListForEach(l *List, f func(*NodeL)) {
	it := l.Head
	for it != nil {
		f(it)
		it = it.Next
	}
}
