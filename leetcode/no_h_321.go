package leetcode

import (
	"fmt"
)

//TODO
func maxNumber321S1(nums1 []int, nums2 []int, k int) []int {
	result := make([]int, k)
	l1, l2 := len(nums1), len(nums2)
	dict := make([][][]int, 2)
	dict[0] = make([][]int, 10)
	dict[1] = make([][]int, 10)
	for i, n := range nums1 {
		dict[0][n] = append(dict[0][n], i)
	}
	for i, n := range nums2 {
		dict[1][n] = append(dict[1][n], i)
	}
	j := k
	anchor1, anchor2 := 0, 0
	fmt.Println(dict[0])
	fmt.Println(dict[1])
	for i := 9; i >= 0; i-- {
		if j == 0 {
			break
		}
		index1, index2 := 0, 0
		fmt.Println("--------------------", i, "--------------------")
		for index1 < len(dict[0][i]) || index2 < len(dict[1][i]) {
			fmt.Println("-----------------", anchor1, anchor2, "-----------------")
			fmt.Println(len(dict[0][i]), dict[0][i])
			fmt.Println(len(dict[1][i]), dict[1][i])
			fmt.Println("当前", dict[1][i][index2], "剩余", j, "NUMS LEFT", l1-anchor1-1, l2-dict[1][i][index2])
			if index1 >= len(dict[0][i]) {
				fmt.Println("B1")
				if (l2-dict[1][i][index2])+l1-anchor1-1 < j {
					break
				}
				if dict[1][i][index2] >= anchor2 {
					result[k-j] = i
					anchor2 = dict[1][i][index2]
					j--
				}
				index2++
			} else if index2 >= len(dict[1][i]) {
				fmt.Println("A1")
				if (l1-dict[0][i][index1])+l2-anchor2-1 < j {
					break
				}
				if dict[0][i][index1] >= anchor1 {
					result[k-j] = i
					anchor1 = dict[0][i][index1]
					j--
				}
				index1++
			} else {
				if l1-dict[0][i][index1] > l2-dict[1][i][index2] {
					fmt.Println("A")
					if (l1-dict[0][i][index1])+l2-anchor2-1 < j {
						break
					}
					if dict[0][i][index1] >= anchor1 {
						result[k-j] = i
						anchor1 = dict[0][i][index1]
						j--
					}
					index1++
				} else {
					fmt.Println("B")
					if (l2-dict[1][i][index2])+l1-anchor1-1 < j {
						break
					}
					if dict[1][i][index2] >= anchor2 {
						result[k-j] = i
						anchor2 = dict[1][i][index2]
						j--
					}
					index2++
				}
			}
		}
	}
	return result
}

//TODO
func maxNumber321S2(nums1 []int, nums2 []int, k int) []int {
	result := make([]int, k)
	// fmt.Println(findList(nums1, k))
	fmt.Println(findList(nums2, k))
	n1 := findList(nums1, k)
	n2 := findList(nums2, k)
	i1, i2 := 0, 0
	index := 0
	for k != index {
		if i1 == len(n1) {
			result[index] = n1[i2]
			i2++
		} else if i2 == len(n2) {
			result[index] = n1[i1]
			i1++
		} else if n1[i1] > n2[i2] {
			result[index] = n1[i1]
			i1++
		} else {
			result[index] = n1[i2]
			i2++
		}
		index++
	}

	return result
}

func findList(nums []int, k int) []int {
	l := len(nums)
	var result []int
	if k > l {
		result = make([]int, l)
		k = l
	} else {
		result = make([]int, k)
	}
	dict := make([][]int, 10)
	for i, n := range nums {
		dict[n] = append(dict[n], i)
	}
	count, anchor := 0, -1

	for count != k {
	a:
		for i := 9; i >= 0; i-- {
			for _, index := range dict[i] {
				if anchor < index {
					if k-count <= l-index {
						result[count] = nums[index]
						count++
						anchor = index
						break a
					} else {
						break
					}
				}
			}
		}
	}

	return result
}
