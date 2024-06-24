package piscine

func ListPushFront(l *List, data interface{}) {
	Node := &NodeL{Data: data, Next: nil}
	Node.Next = l.Head
	l.Head = Node
}
