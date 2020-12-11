package leetcode

//DONE
func leastIntervalS1(tasks []byte, n int) int {
	m := make([]int, 26)
	dict := make([]int, 26)
	for _, task := range tasks {
		m[int(task-'A')]++
		dict[int(task-'A')] = 1
	}
	for i := 25; i > 0; i-- {
		if dict[i] == 0 {
			dict = append(dict[:i], dict[i+1:]...)
			m = append(m[:i], m[i+1:]...)
		}
	}

	result, count := 0, len(tasks)
	for count != 0 {
		max, index := 0, -1
		for i, schedule := range dict {
			if schedule == 1 {
				if m[i] > max {
					max = m[i]
					index = i
				}
			} else {
				if schedule == n+1 {
					dict[i] = 1
				} else {
					dict[i]++
				}
			}
		}
		if index != -1 {
			m[index]--
			if m[index] == 0 {
				dict = append(dict[:index], dict[index+1:]...)
				m = append(m[:index], m[index+1:]...)
			} else {
				if dict[index] == n+1 {
					dict[index] = 1
				} else {
					dict[index]++
				}
			}
			count--
		}
		result++
	}

	return result
}

//TODO 构造方式
func leastIntervalS2(tasks []byte, n int) int {
	return 0
}
