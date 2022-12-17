package piscine

func ConcatParams(args []string) string {
	// if len(args) < 0 {
	// 	return ""
	// }
	// l := 0

	// for i := 0; i < len(args); i++ {
	// 	l += len(args[i]) + 1
	// }
	// l -= 1
	// answer := make([]rune, l)

	// st := 0
	// i := 0
	// for j := 0; j < len(args); j++ {

	// 	end := st + len(args[j])
	// 	for k := st; k < end; k++ {
	// 		answer[k] = rune(args[j][i])
	// 		i++
	// 		// for i := 0; i < l; i++ {
	// 		// 	if k == len(args[j])-1 && j != len(args)-1  && l{
	// 		// 		answer[i+1] = rune('\n')
	// 		// 	} else {
	// 		// 		answer[i] = rune(args[j][k])
	// 		// 	}
	// 		// }
	// 	}
	// 	i = 0
	// 	st = end + 1
	// 	if j != len(args)-1 {
	// 		answer[st] = rune('\n')
	// 		st += 1
	// 	}

	// }

	answer := ""
	for i := 0; i < len(args); i++ {
		answer += args[i]
		if i != len(args)-1 {
			answer += "\n"
		}
	}

	return string(answer)
}
