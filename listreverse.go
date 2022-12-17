package piscine

func ListReverse(l *List) {
	var prev *NodeL
	cur := l.Head
	var next *NodeL

	for cur != nil {

		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
}
