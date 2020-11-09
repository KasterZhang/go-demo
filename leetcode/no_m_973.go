package leetcode

func kClosest973S1(points [][]int, K int) [][]int {
	result := [][]int{}
	points = qs(points)
	for K > 0 {
		K--
		result = append(result, points[K])
	}
	return result
}

func qs(points [][]int) [][]int {
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
	return append(left, right...)
}

func kClosest973S2(points [][]int, K int) [][]int {
	result := [][]int{}
	return result
}
