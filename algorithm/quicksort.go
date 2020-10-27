package algorithm

func quickSort(source [][]int) [][]int {
	if len(source) <= 1 {
		return source
	}
	var base []int
	left := [][]int{}
	right := [][]int{}
	for i, item := range source {
		if i == 0 {
			base = item
		} else {
			if item[1] < base[1] {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
	}
	return append(append(quickSort(left), base), quickSort(right)...)
}
