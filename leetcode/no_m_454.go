package leetcode

//DONE
func fourSumCount454(A []int, B []int, C []int, D []int) int {
	dm := make(map[int]int, len(A)*len(B))
	for _, n := range A {
		for _, m := range B {
			if count, ok := dm[n+m]; ok {
				dm[n+m] = count + 1
			} else {
				dm[n+m] = 1
			}
		}
	}
	result := 0
	for _, n := range C {
		for _, m := range D {
			if count, ok := dm[-n-m]; ok {
				result += count
			}
		}
	}
	return result
}
