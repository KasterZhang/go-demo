package leetcode

//TODO
func allCellsDistOrder1030S1(R int, C int, r0 int, c0 int) [][]int {
	side := 1
	result := make([][]int, R*C)
	board := make([][]int, R)
	for i := range board {
		line := make([]int, C)
		board[i] = line
	}
	result[0] = []int{r0, c0}
	board[r0][c0] = 1
	i := 1
	for i < R*C {
		for j := 0; j <= side; j++ {
			if r0+side < R && c0+j < C && board[r0+side][c0+j] == 0 {
				board[r0+side][c0+j] = 1
				result[i] = []int{r0 + side, c0 + j}
				i++
			}
			if r0+side < R && c0-j >= 0 && board[r0+side][c0-j] == 0 {
				board[r0+side][c0-j] = 1
				result[i] = []int{r0 + side, c0 - j}
				i++
			}
			if r0-side >= 0 && c0+j < C && board[r0-side][c0+j] == 0 {
				board[r0-side][c0+j] = 1
				result[i] = []int{r0 - side, c0 + j}
				i++
			}
			if r0-side >= 0 && c0-j >= 0 && board[r0-side][c0-j] == 0 {
				board[r0-side][c0-j] = 1
				result[i] = []int{r0 - side, c0 - j}
				i++
			}
			if c0+side < C && r0+j < R && board[r0+j][c0+side] == 0 {
				board[r0+j][c0+side] = 1
				result[i] = []int{r0 + j, c0 + side}
				i++
			}
			if c0+side < C && r0-j >= 0 && board[r0-j][c0+side] == 0 {
				board[r0-j][c0+side] = 1
				result[i] = []int{r0 - j, c0 + side}
				i++
			}
			if c0-side >= 0 && r0+j < R && board[r0+j][c0-side] == 0 {
				board[r0+j][c0-side] = 1
				result[i] = []int{r0 + j, c0 - side}
				i++
			}
			if c0-side >= 0 && r0-j >= 0 && board[r0-j][c0-side] == 0 {
				board[r0-j][c0-side] = 1
				result[i] = []int{r0 - j, c0 - side}
				i++
			}
		}
		side++
	}
	return result
}

func allCellsDistOrder1030S2(R int, C int, r0 int, c0 int) [][]int {
	board := make([][]int, R)
	for i := range board {
		line := make([]int, C)
		board[i] = line
	}
	side := 1
	result := make([][]int, R*C)
	result[0] = []int{r0, c0}
	i := 1
	for i != R*C-1 {
		tempR, tempC := side, 0
		for tempR >= 0 {
			if r0-tempR >= 0 && c0-tempC >= 0 {
				result[i] = []int{r0 - tempR, c0 - tempC}
				i++
			}
			if r0-tempR >= 0 && c0+tempC < C {
				result[i] = []int{r0 - tempR, c0 + tempC}
				i++
			}
			if r0+tempR < R && c0-tempC >= 0 {
				result[i] = []int{r0 + tempR, c0 - tempC}
				i++
			}
			if r0+tempR < R && c0+tempC < C {
				result[i] = []int{r0 + tempR, c0 + tempC}
				i++
			}
			tempR--
			tempC++
		}
	}
	return result
}
