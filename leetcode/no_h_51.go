package leetcode

import "strings"

//DONE
//IDEA 通过旋转来优化
func solveNQueens51S1(n int) [][]string {
	checkerboard := make([]string, n, n)
	var line strings.Builder
	for i := 0; i < n; i++ {
		line.WriteString("-")
	}
	for i := range checkerboard {
		checkerboard[i] = line.String()
	}
	result := [][]string{}
	for i := 0; i < n; i++ {
		result = append(result, markBoard(0, i, n, checkerboard)...)
	}
	// for _,board := range result {
	//     for _, row := range board {
	//         fmt.Println(row)
	//     }
	//     fmt.Println("*****************")
	// }
	return result
}

func markBoard(r, c, n int, temp []string) [][]string {
	board := make([]string, n)
	copy(board, temp)
	result := make([][]string, 0, n)
	rstr := []byte(board[r])
	rstr[c] = 'Q'
	board[r] = string(rstr)
	if r+1 == len(board) {
		result = append(result, board)
		return result
	}
	for rowNo, rstr := range board {
		rowBuffer := abs(rowNo - r)
		row := []byte(rstr)
		if rowNo == r {
			for colNo := range row {
				if colNo != c {
					row[colNo] = '.'
				}
			}
		} else {
			for colNo := range row {
				colBuffer := abs(colNo - c)
				if colNo == c {
					row[colNo] = '.'
				}
				if rowBuffer == colBuffer {
					row[colNo] = '.'
				}
			}
		}
		board[rowNo] = string(row)
	}
	if r+1 < n {
		availableNode := false
		for colNo, col := range board[r+1] {
			if col != '.' {
				availableNode = true
				rs := markBoard(r+1, colNo, n, board)
				if rs != nil {
					result = append(result, rs...)
				}
			}
		}
		if !availableNode {
			return nil
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
