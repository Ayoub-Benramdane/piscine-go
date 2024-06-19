package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
