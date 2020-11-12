package leetcode

func kClosest973S1(points [][]int, K int) [][]int {
	return qs973S1(points, K)
}

func qs973S1(points [][]int, count int) [][]int {
	if len(points) <= 1 {
		return points
	}
	base := points[0][0]*points[0][0] + points[1][1]*points[1][1]
	left, right := [][]int{}, [][]int{}
	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		if distance <= base {
			left = append(left, point)
		} else {
			right = append(right, point)
		}
	}
	if count == len(left) {
		return left
	} else if count > len(left) {
		count -= len(left)
		count--
		if count == 0 {
			return append(left, points[0])
		}
		return append(append(left, points[0]), qs973S1(right, count)...)
	} else if count < len(left) {
		return qs973S1(left, count)
	}
	return append(left, right...)
}
