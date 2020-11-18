package leetcode

import "fmt"

//TODO
func getMaxGridHappiness1659(m int, n int, introvertsCount int, extrovertsCount int) int {
	board := make([][]int, m)
	for i := range board {
		line := make([]int, n)
		board[i] = line
	}
	result := 0
	preIntroverts := 0
	fmt.Println(introvertsCount, extrovertsCount, board)
a:
	for introvertsCount != 0 {
		for i, line := range board {
			for j, cell := range line {
				if cell == 0 {
					intro, extro := check(i, j, board)
					fmt.Println(intro, extro)
					if (intro + extro) == preIntroverts {
						board[i][j] = 1
						introvertsCount--
						fmt.Println("introvertsCount", introvertsCount)
						result += 120 - preIntroverts*30
						if introvertsCount == 0 {
							break a
						}
					}
				}
			}
		}
		fmt.Println("preIntroverts", preIntroverts)
		preIntroverts++
	}
	fmt.Println(board)
	index := 0
b:
	for extrovertsCount != 0 {
	c:
		for index != 4 {
			temp := 4 - index
			for temp >= 0 {
				for i, line := range board {
					for j, cell := range line {
						if cell == 0 {
							fmt.Println("temp", temp, "index", index)
							intro, extro := check(i, j, board)
							if extro == temp && intro == index {
								index = 0
								board[i][j] = 2
								extrovertsCount--
								fmt.Println("extrovertsCount", extrovertsCount)
								result += 40 + 20*extro - 30*intro
								if extrovertsCount == 0 {
									break b
								}
								break c
							}
						}
					}
				}
				temp--
			}
			index++
		}
		fmt.Println("extrovertsCount", extrovertsCount)
	}
	fmt.Println(board)
	return result
}

func check(i, j int, board [][]int) (int, int) {
	introverts, extroverts := 0, 0
	if i > 0 && board[i-1][j] != 0 {
		if board[i-1][j] == 1 {
			introverts++
		} else {
			extroverts++
		}
	}

	if i < len(board)-1 && board[i+1][j] != 0 {
		if board[i+1][j] == 1 {
			introverts++
		} else {
			extroverts++
		}
	}

	if j > 0 && board[i][j-1] != 0 {
		if board[i][j-1] == 1 {
			introverts++
		} else {
			extroverts++
		}
	}

	if j < len(board[i])-1 && board[i][j+1] != 0 {
		if board[i][j+1] == 1 {
			introverts++
		} else {
			extroverts++
		}
	}
	return introverts, extroverts
}
