package leetcode

//DONE
func minimumJumps1654(forbidden []int, a int, b int, x int) int {
	dp := []int{0}
	count := 0
	if 0 == x {
		return count
	}
	for {
		count++
		move := false
		nextDP := make([]int, 0, len(dp)*2)

		for _, n := range dp {
			if n > 0 {
				if n-b >= 0 && check1654(forbidden, n-b) {
					forbidden = append(forbidden, n-b)
					move = true
					nextDP = append(nextDP, b-n)
					if n-b == x {
						return count
					}
				}

			} else {
				n = -n
			}
			if n+a < x+10*b && check1654(forbidden, n+a) {
				forbidden = append(forbidden, n+a)
				move = true
				nextDP = append(nextDP, n+a)
				if n+a == x {
					return count
				}
			}
		}
		dp = nextDP
		if !move {
			break
		}
	}
	return -1
}

func check1654(forbidden []int, target int) bool {
	for _, n := range forbidden {
		if n == target {
			return false
		}
	}
	return true
}
