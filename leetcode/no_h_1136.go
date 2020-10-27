package leetcode

//DONE
func minimumSemesters1136S1(N int, relations [][]int) int {
	term := 0
	out := make([][]int, N)
	in := make([]int, N)
	for _, road := range relations {
		in[road[1]-1]++
		out[road[0]-1] = append(out[road[0]-1], road[1]-1)
	}

	n := N
	for {
		if n == 0 {
			break
		}
		temp := make([]int, N)
		copy(temp, in)
		find := false
		for i, du := range in {
			if du == 0 {
				find = true
				n--
				temp[i] = -1
				for _, outNode := range out[i] {
					if temp[outNode] <= 0 {
						return -1
					}
					temp[outNode]--
				}
				out[i] = nil
			}
		}
		if !find {
			return -1
		}
		in = temp
		term++
	}
	return term
}
