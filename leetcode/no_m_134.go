package leetcode

import "fmt"

//gas  = [1,2,3,4,5]
// cost = [3,4,5,1,2]

// 输出: 3
//DONE O(N^2) O(1)
func canCompleteCircuit134S1(gas []int, cost []int) int {
	for i, petrol := range gas {
		count := 1
		tank := petrol - cost[i]

		for count < len(cost) && tank >= 0 {
			target := i + count
			if i+count > len(cost)-1 {
				target = target - len(cost)
			}
			tank += gas[target]
			fmt.Println("location", i, "target", target, "tank", tank)
			tank -= cost[target]
			if tank < 0 {
				break
			}
			count++
		}
		if tank >= 0 && count == len(cost) {
			return i
		}
		fmt.Println("--------")
	}

	return -1
}

//DONE O(N) O(1)
func canCompleteCircuit134S2(gas []int, cost []int) int {
	i, c := 0, 0
a:
	for i < len(gas) && c < len(gas) {
		c++
		count := 1
		tank := gas[i] - cost[i]

		for count < len(cost) && tank >= 0 {
			target := i + count
			if i+count > len(cost)-1 {
				target = target - len(cost)
			}
			tank += gas[target]
			tank -= cost[target]
			if tank < 0 {
				i = target
				continue a
			}
			count++
		}
		if tank >= 0 && count == len(cost) {
			return i
		}
		i++
	}

	return -1
}
