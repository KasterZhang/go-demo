package leetcode

//DONE
func decrypt1652(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}
	result := make([]int, len(code))
	if k > 0 {
		for i := range result {
			count, sum := 0, 0
			for count < k {
				if i+1+count <= len(code)-1 {
					sum += code[i+1+count]
				} else {
					sum += code[i+1+count-len(code)]
				}
				count++
			}
			result[i] = sum
		}
	}

	if k < 0 {
		for i := range result {
			count, sum := 0, 0
			for count < -k {
				if i-count-1 >= 0 {
					sum += code[i-1-count]
				} else {
					sum += code[len(code)-1+i-count]
				}
				count++
			}
			result[i] = sum
		}
	}
	return result
}
