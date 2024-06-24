package piscine

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	it := l.Head
	for it != nil {
		if comp(it.Data, ref) {
			return &it.Data
		}
		it = it.Next
	}
	return &it.Data
}
