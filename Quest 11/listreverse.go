package piscine

func ListReverse(l *List) {
	var previous, current *NodeL
	current = l.Head
	l.Tail = l.Head
	for current != nil {
		previous, current, current.Next = current, current.Next, previous
	}
	l.Head = previous
}
