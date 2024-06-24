package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if cond(it) {
			f(it)
		}
		it = it.Next
	}
}
