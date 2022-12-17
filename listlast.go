package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else if l.Tail == nil {
		return l.Head.Data
	} else {
		return l.Tail.Data
	}
}
