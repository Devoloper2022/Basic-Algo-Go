package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	cur := l.Head
	// var prev NodeL

	for cur != nil {
		if l.Head == data_ref {
			cur = l.Head.Next
			l.Head = cur
		} else if cur == data_ref {
		}
	}
}
