package piscine

func ListMerge(l1 *List, l2 *List) {
	it := l2.Head
	for it != nil {
		ListPushBack(l1, it.Data)
		it = it.Next
	}
}
