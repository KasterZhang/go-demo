package leetcode

//DONE TODO 双BFS
func ladderLength127(beginWord string, endWord string, wordList []string) int {
	queue := []string{beginWord}
	result := 2
	dict := wordList
	for len(queue) != 0 {
		l := len(queue)
		for l != 0 {
			pop := queue[0]
			queue = queue[1:]
			for i, word := range dict {
				if word != "" {
					if judge127(pop, word) {
						dict[i] = ""
						queue = append(queue, word)
						if word == endWord {
							return result
						}
					} else if pop == word {
						dict[i] = ""
					}
				}
			}
			l--
		}
		result++
	}

	return 0
}

func judge127(a, b string) bool {
	for i := range a {
		if a[i] != b[i] {
			if a[i+1:] == b[i+1:] {
				return true
			}
			break
		}
	}
	return false
}
