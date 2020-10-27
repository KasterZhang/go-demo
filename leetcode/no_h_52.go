package leetcode

import "fmt"

//TODO
func totalNQueens52(n int) int {
	checkboard := make([][]rune, n, n)
	for i := range checkboard {
		checkboard[i] = make([]rune, n)
	}
	count := 0
	for i := 0; i < n; i++ {
		count += mark52(0, i, checkboard)
	}
	return count
}

func mark52(r, c int, checkboard [][]rune) int {
	n := len(checkboard)
	board := make([][]rune, n, n)
	for i := range board {
		board[i] = make([]rune, n)
		copy(board[i], checkboard[i])
	}
	for _, row := range board {
		fmt.Println(row)
	}
	if r == n-1 {
		return 1
	}
	fmt.Printf("-----%d %d------\n", r, c)
	for i, row := range board {
		for j := range row {
			if i == r && j == c {
				board[i][j] = 2
			} else if j == c {
				board[i][c] = 1
			} else if i == r {
				board[r][j] = 1
			} else if r-i == c-j || r-i == j-c {
				board[i][j] = 1
			}
		}
	}
	// for _,row := range board {
	//         fmt.Println(row)
	// }
	fmt.Println("--------------")
	r++
	count := 0
	for i, col := range board[r] {
		if col == 0 {
			count += mark52(r, i, board)
		}
	}

	return count
}
