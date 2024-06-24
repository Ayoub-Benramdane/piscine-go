package main

import (
	"fmt"
)

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}

type NodeI struct {
	Data int
	Next *NodeI
}

var V = 0

func ListSort(l *NodeI) *NodeI {
	it := l
	for it.Next != nil {
		if it.Data > it.Next.Data {
			it, it.Next = it.Next, it
			fmt.Println(it.Data, it.Next.Data)
			it = it.Next
			fmt.Println(it.Data, it.Next.Data)
			ListSort(it)
		}
	}
	return l
}
