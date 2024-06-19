package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	Node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = Node
		l.Tail = Node
	} else {
		l.Tail.Next = Node
		l.Tail = Node
	}
}
