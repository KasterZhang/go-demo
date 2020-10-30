package leetcode

func islandPerimeter(grid [][]int) int {
	result := 0

	//从2个角度记述
	for i, row := range grid {
		for j, g := range row {
			if g == 1 {
				if i > 0 {
					if grid[i-1][j] != 1 {
						result++
					}
				} else {
					result++
				}
				if i < len(grid)-1 {
					if grid[i+1][j] != 1 {
						result++
					}
				} else {
					result++
				}

				if j > 0 {
					if grid[i][j-1] != 1 {
						result++
					}
				} else {
					result++
				}
				if j < len(row)-1 {
					if grid[i][j+1] != 1 {
						result++
					}
				} else {
					result++
				}
			}
		}
	}
	return result
}
