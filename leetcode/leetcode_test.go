package leetcode

import (
	"fmt"
	"testing"
)

func TestGame(*testing.T) {
	// edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {2, 5}, {5, 6}}
	// fmt.Println(chaseGame(edges, 3, 5))

	// edges := [][]int{{1, 2}, {2, 3}, {3, 1}, {3, 6}, {2, 4}, {4, 5}, {5, 8}, {4, 7}}
	// fmt.Println(chaseGame(edges, 6, 7))
}

func Test78(*testing.T) {
	// r := []int{5, 2, 3, 1, 7, 4, 6}
	r := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 0}
	fmt.Println(subsets(r))
}

func TestDemo(*testing.T) {
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:2]
	fmt.Println("slice1:", append(slice1, 6, 7, 8))
	fmt.Println("arr1:", arr1)
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:3]

	fmt.Println("slice2:", append(slice2, 6, 7, 8))
	fmt.Println("arr2:", arr2)
}

func Test51(*testing.T) {
	// solveNQueens(5)
	solveNQueens51S1(7)
}

func Test106(*testing.T) {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}

func Test37(*testing.T) {
	var board [][]byte
	astr := [][]string{
		[]string{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		[]string{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		[]string{".", "9", "8", ".", ".", ".", ".", "6", "."},
		[]string{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		[]string{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		[]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		[]string{".", "6", ".", ".", ".", ".", "2", "8", "."},
		[]string{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		[]string{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	for _, arr := range astr {
		tmp := []byte{}
		for _, ch := range arr {
			tmp = append(tmp, byte(ch[0]))
		}
		board = append(board, tmp)
	}
	// solveNQueens(5)
	solveSudoku37(board)
}
func Test1136(*testing.T) {
	astr := [][]int{[]int{1, 3}, []int{2, 3}}
	fmt.Println(minimumSemesters1136S1(3, astr))
}

func Test42(*testing.T) {
	// arr1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	arr2 := []int{4, 2, 0, 3, 2, 5}
	// fmt.Println(trap42(arr1)) //6
	fmt.Println(trap42(arr2)) //9
}

func Test4(*testing.T) {
	//[1,3,5,7,9,11,13,15,17,19]
	//[2,4,6,8,10]
	nums1 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	nums2 := []int{2, 4, 6, 8, 10}
	fmt.Println(findMedianSortedArrays4S1(nums1, nums2))
}
