package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	st := l.Head
	for st != nil {
		if comp(st.Data, ref) {
			return &st.Data
		}

		st = st.Next
	}
	return nil
}
