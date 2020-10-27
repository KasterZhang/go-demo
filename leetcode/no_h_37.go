package leetcode

import "bytes"

//DONE
func solveSudoku37(board [][]byte) {
	//49-57
	sovle37(0, 0, board)
}
func sovle37(s_x, s_y int, board [][]byte) bool {
	for y, row := range board {
		if y >= s_y {
			for x, ch := range row {
				if ((y == s_y && x >= s_x) || y > s_y) && ch == 46 {
					avilible := getLeagle37(x, y, board)
					if len(avilible) == 0 {
						return false
					}
					for _, b := range avilible {
						board[y][x] = b
						// fmt.Println("位置 x:", x, "y:", y, "所有可能值:", string(avilible), "插入可能值:", string(b))
						if sovle37(x, y, board) {
							return true
						}

						board[y][x] = 46

					}
					return false
				}
			}
		}
	}
	return true
}

func getLeagle37(x, y int, board [][]byte) []byte {
	tmp := []byte{49, 50, 51, 52, 53, 54, 55, 56, 57}
	//行
	for _, b := range board[y] {
		if b != 46 {
			index := bytes.IndexByte(tmp, b)
			if index != -1 {
				tmp = append(tmp[:index], tmp[index+1:]...)
			}
		}
	}
	//列
	for _, bs := range board {
		if bs[x] != 46 {
			index := bytes.IndexByte(tmp, bs[x])
			if index != -1 {
				tmp = append(tmp[:index], tmp[index+1:]...)
			}
		}
	}
	//域
	for i := x / 3 * 3; i <= x/3*3+2; i++ {
		for j := y / 3 * 3; j <= y/3*3+2; j++ {
			if board[j][i] != 46 {
				index := bytes.IndexByte(tmp, board[j][i])
				if index != -1 {
					tmp = append(tmp[:index], tmp[index+1:]...)
				}
			}
		}
	}
	return tmp
}
