package leetcode

//TODO DP 解法
//DONE
func findRotateSteps(ring string, key string) int {
	ll, min := len(key), len(key)+len(key)*len(key)

	m := step(ring, rune(key[0]))
	for loc, r := range key[1:] {
		tempM := make(map[string]int)
		for t, s := range m {
			for tt, ss := range step(t, r) {
				if loc+2 == ll {
					if ss+s < min {
						min = ss + s
					}
				} else {
					if pre, ok := tempM[tt]; ok {
						if ss+s < pre {
							tempM[tt] = ss + s
						}
					} else {
						tempM[tt] = ss + s
					}
				}

			}
		}
		m = tempM
	}

	return min + ll
}

func step(ring string, target rune) map[string]int {
	result, last := make(map[string]int), len(ring)-1
	for i, r := range ring {
		if i > (last+1)/2 {
			break
		}
		if r == target {
			if _, ok := result[ring[:i]+ring[i:]]; !ok {
				result[ring[i:]+ring[:i]] = i
			}
		}
		if rune(ring[last-i]) == target {
			if _, ok := result[ring[last-i:]+ring[:last-i]]; !ok {
				result[ring[last-i:]+ring[:last-i]] = i + 1
			}
		}
	}
	return result
}
