package leetcode

import "sort"

//DONE
func sortByBits1356(arr []int) []int {
	// arr = qs(arr)
	sort.Ints(arr)
	dict := make(map[int][]int)
	keys := make([]int, 0, len(arr))
	for i, n := range arr {
		c := 0
		for n != 0 {
			n &= n - 1
			c++
		}
		if _, ok := dict[c]; ok {
			dict[c] = append(dict[c], arr[i])
		} else {
			keys = append(keys, c)
			dict[c] = []int{arr[i]}
		}
	}
	sort.Ints(keys)
	// keys = qs(keys)
	j := 0
	for _, key := range keys {
		for _, n := range dict[key] {
			arr[j] = n
			j++
		}
	}
	return arr
}
