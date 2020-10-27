package leetcode

import "fmt"

//题目:https://leetcode-cn.com/problems/Za25hA/
//分析:
//首先分析出图中的环,寻找大于3的环(3的环会被追上)
//Hunter:
//因为是猎人先追,所以猎人的策略就是寻找到猎物的最短路径并移动一步
//Prey:
//如果存在3以上的环则判断所有环中节点,自己是否可以先于猎人进入环
//中猎物遍历环中所有点获取路径,同时获取猎人从当前位置到达目标点的
//所有路径,如果所差步数大于等于2步,则可以顺利逃致环中 返回-1否则
//,反之不能可以逃入环中.倒叙遍历猎人到当前所有点路径如果自己到达
//目标点位的步数比猎人快至少2步则向目标点移动.返回猎人到达目标点
//所需步数
func chaseGame(edges [][]int, hunter int, prey int) int {
	preHunterRoads, _ := anaylse(edges, hunter)
	hunter = preHunterRoads[prey-1][1]
	if hunter == prey {
		return 1
	}
	hunterRoads, ring := anaylse(edges, hunter)
	preyRoads, _ := anaylse(edges, prey)
	if len(ring) > 0 {
		for ringPoint, _ := range ring {
			//以为hunter已经移动过了 所以+1
			if len(preyRoads[ringPoint-1])+1 <= len(hunterRoads[ringPoint-1]) {
				return -1
			}
		}
	}

	maxStep := 0
	for i, roads := range hunterRoads {
		if maxStep < len(roads) && len(preyRoads[i])+1 <= len(roads) {
			maxStep = len(roads)
		}
	}
	fmt.Println("hunter ", hunter)
	fmt.Println("prey ", prey)
	fmt.Println("ring ", preHunterRoads)
	fmt.Println("preHunterRoads ", preHunterRoads)
	fmt.Println("hunterRoads ", hunterRoads)
	fmt.Println("preyRoads ", preyRoads)
	fmt.Println("maxStep ", maxStep)
	return maxStep
}

//寻找环
//如何找大于3个点的环呢DFS BFS应该都可以
//代表到达目标点位的最佳路径,顺便找到环
func anaylse(edges [][]int, from int) ([][]int, map[int]struct{}) {
	fmt.Println(" ========================== ")
	ring := map[int]struct{}{}
	best := make([][]int, len(edges))
	best[from-1] = []int{from}
	ringArr := make([]int, 0, len(edges))
	var nextCh chan []int = make(chan []int, len(edges))
	nextCh <- []int{from}
	defer close(nextCh)
	//已标记的点
	markPoints := make([]int, 0, len(edges))
	isClose := false
	// loop:
	for {
		source, _ := <-nextCh
		if isClose {
			fmt.Println("通道关闭!! ")
			break
		}
		//取尾节点
		start := source[len(source)-1]
		if !has(markPoints, start) {
			markPoints = append(markPoints, start)
		}
		// 获取尾节点可以接触的下一组节点
		points, hasNext := next(edges, start, markPoints)
		// fmt.Println("chan取出资源 ", source)
		fmt.Println(" ------------------------- ")
		fmt.Println("通道是否关闭 ", isClose)
		fmt.Println("当前节点 ", start)
		fmt.Println("已标记节点 ", markPoints)
		fmt.Println(hasNext, " 获取尾节点可以接触的下一组节点 ", points)
		if hasNext {
			//p代表下一个节点
			for _, np := range points {
				fmt.Println("下一个节点 ", np)
				tmp := append(source, np)
				//找到了大于3的环
				//下一个节点已存在最优路径 并且 下一个节点最优路径的末尾等于下一个节点 =>遭遇可能的环状事件
				if len(best[np-1]) > 0 && best[np-1][len(best[np-1])-1] == np {
					ring = getRing(tmp, best[np-1], ring)
					// fmt.Println("ringArr ", ringArr)
					// fmt.Println("tmp ", tmp)
					// fmt.Println("best[p-1] ", best[np-1])
					// fmt.Println("removeDuplicateElement ", removeDuplicateElement(append(tmp, best[np-1]...)...))
					// ringArr = append(ringArr, removeDuplicateElement(append(tmp, best[np-1]...)...)...)
				} else {
					best[np-1] = append(best[np-1], tmp...)
					fmt.Println(np, "点记录最佳路径 ", best[np-1])
				}
				nextCh <- tmp
			}
		}
		fmt.Println("已标记节点 ", len(markPoints), "边数 ", len(edges))
		if len(markPoints) == len(edges) {
			isClose = true
		}
	}
	for _, point := range ringArr {
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

func getRing(a, b []int, ring map[int]struct{}) map[int]struct{} {
	base := []int{}
	line := []int{}
	result := []int{}
	a = clean(a, ring)
	b = clean(b, ring)
	if len(a) < len(b) {
		base = a
		line = b
	} else {
		base = b
		line = a
	}

	count := 0
	tempRing := []int{}
	for i, v := range base {
		if v == line[i] && count == 0 {
			// fmt.Println("重复元素 抛弃")
		} else if v == line[i] && count == 1 {
			tempRing = []int{}
			count = 0
		} else if v == line[i] && count > 1 {
			tempRing = append(tempRing, v)
			fmt.Println("找到环 记录 ", tempRing)
			result = append(result, tempRing...)
			tempRing = []int{}
			count = 0
		} else {
			count++
			tempRing = append(tempRing, base[i-1], v)
		}
	}
	for _, point := range result {
		ring[point] = struct{}{}
	}
	return ring
}

func clean(arr []int, without map[int]struct{}) []int {
	aa := []int{}
	for _, v := range arr {
		if _, ok := without[v]; !ok {
			aa = append(aa, v)
		}
	}
	return aa
}
