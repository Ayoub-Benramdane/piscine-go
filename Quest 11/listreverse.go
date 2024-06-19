package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var prev, current *NodeL
	current = l.Head
	l.Tail = l.Head
	for current != nil {
		prev, current, current.Next = current, current.Next, prev
	}
	l.Head = prev
}
