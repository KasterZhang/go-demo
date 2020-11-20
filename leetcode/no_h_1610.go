package leetcode

import (
	"fmt"
	"math"
	"sort"
)

//TODO
func visiblePoints1610(points [][]int, angle int, location []int) int {
	for i, p := range points {
		points[i] = []int{p[0] - location[0], p[1] - location[1]}
	}

	angles := make([]float64, 0, len(points))
	yuandian := 0
	for _, n := range points {
		x, y := n[0], n[1]
		if x == 0 && y == 0 {
			yuandian++
		} else {
			angles = append(angles, math.Atan2(float64(y), float64(x))*180/math.Pi)
		}
	}

	sort.Float64s(angles)
	result := yuandian

	//双指针滑动窗
	for _, a := range angles {
		min, max := a, a+float64(angle)
		temp := yuandian
		for j, t := range angles {
			if (max >= 180 && (min <= t && t <= 180 || -180 <= t && t <= max-360)) || (min <= t && t <= max) {
				fmt.Println("j", j)
				temp++
			}
		}
		if temp > result {
			result = temp
		}
	}
	return result
}
