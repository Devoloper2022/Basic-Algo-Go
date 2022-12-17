package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = node
	} else if l.Tail == nil {
		l.Tail = l.Head
		l.Head = node
		l.Head.Next = l.Tail
	} else {
		node.Next = l.Head
		l.Head = node
	}
}
