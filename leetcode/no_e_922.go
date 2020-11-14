package leetcode

//DONE
func sortArrayByParityII922S1(A []int) []int {
	odd, even, result := make([]int, 0, len(A)/2), make([]int, 0, len(A)/2), make([]int, len(A))
	for _, n := range A {
		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}
	for i := range result {
		if i%2 == 0 {
			result[i] = even[i/2]
		} else {
			result[i] = odd[(i-1)/2]
		}
	}
	return result
}

func sortArrayByParityII922S2(A []int) []int {
	result := make([]int, len(A))
	even, odd := 0, 1
	for _, n := range A {
		if n%2 == 0 {
			result[even] = n
			even += 2
		} else {
			result[odd] = n
			odd += 2
		}
	}

	return result
}
