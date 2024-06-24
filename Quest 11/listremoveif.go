package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	it := l.Head
	l.Head = nil
	l.Tail = nil
	for it != nil {
		if it.Data == data_ref {
			it = it.Next
		} else {
			ListPushBack(l, it.Data)
			it = it.Next
		}
	}
}
