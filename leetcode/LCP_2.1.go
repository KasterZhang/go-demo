package leetcode

import "fmt"

var points []int = []int{1, 2, 3, 4, 5, 6}
var edges [][]int = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {2, 5}, {5, 6}}

func chaseGame(edges [][]int, startA int, startB int) int {
	//设定A为追踪者 B为逃跑者
	roads, ring := anaylse(edges, startA)
	fmt.Println(roads, ring)

	return -1
}

//遍历所有可走路径,判断距离prey的最佳路径
func trace(edges [][]int, from int, prey int) (int, bool) {
	return -1, false
}

//遍历所有可走路径,判断距离hunter的距离是否变远,如果全部未变远则放弃行动,否则行动
//趋向于向环移动
//如果能逃入环中则游戏结束 返回-1
//return target move? over?
func escape(edges [][]int, ring map[int]struct{}, from int, hunter int, distance int) (int, bool, bool) {
	points, hasNext := next(edges, from, []int{})
	if hasNext {
		for _, point := range points {
			if _, ok := ring[point]; ok {
				return -1, false, false
			}
		}
	}
	return -1, false, false
}

//代表到达目标点位的最佳路径
func anaylse(edges [][]int, from int) ([][]int, map[int]struct{}) {
	best := make([][]int, len(edges))
	best[from-1] = []int{from}
	temp := make([]int, 0, len(edges))
	var nextCh chan []int = make(chan []int, len(edges))
	nextCh <- []int{from}
	markPoints := []int{}
	for {
		source, ok := <-nextCh
		if !ok {
			break
		}
		start := source[len(source)-1]
		markPoints = append(markPoints, start)
		points, hasNext := next(edges, start, markPoints)
		if hasNext {
			for _, p := range points {
				tmp := append(source, p)
				if len(best[p-1]) > 0 && best[p-1][len(best[p-1])-1] == p {
					fmt.Println(removeDuplicateElement(append(tmp, best[p-1]...)...))
					temp = append(temp, removeDuplicateElement(append(tmp, best[p-1]...)...)...)
				} else {
					best[p-1] = append(best[p-1], tmp...)
				}
				nextCh <- tmp
			}
		}
		if len(markPoints) == len(edges) {
			close(nextCh)
		}
	}
	ring := map[int]struct{}{}
	for _, point := range temp {
		ring[point] = struct{}{}
	}
	return best, ring
}

func next(edges [][]int, location int, without []int) ([]int, bool) {
	nextLocation := []int{}
	for _, v := range edges {
		if v[0] == location && !has(without, v[1]) {
			nextLocation = append(nextLocation, v[1])
		} else if v[1] == location && !has(without, v[0]) {
			nextLocation = append(nextLocation, v[0])
		}
	}
	return nextLocation, len(nextLocation) > 0
}

func has(points []int, point int) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

func removeDuplicateElement(points ...int) []int {
	result := make([]int, 0, len(points))
	temp := map[int]struct{}{}
	for _, point := range points {
		if _, ok := temp[point]; !ok {
			temp[point] = struct{}{}
			result = append(result, point)
		}
	}
	return result
}
