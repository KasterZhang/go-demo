package leetcode

func subsets(nums []int) [][]int {
	//range times equal nums length
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	if len(nums) == 1 {
		return [][]int{nums, []int{}}
	}
	// result := make([][]int, 0, len(nums)*len(nums))
	result := [][]int{}
	// for _, cc := range nums {
	// 	result = append(result, []int{cc})
	// }
	// fmt.Println(result)
	l := len(nums)
	// for i, _ := range nums {
	// 	fmt.Printf("------------- range %d -------------\n", i)
	// 	for ii, _ := range nums {
	// 		if ii > i {
	// 			tmpI := 0
	// 			for iii, ccc := range nums {
	// 				if iii >= ii {
	// 					fmt.Printf("cc %d, source %d, len %d , loc %d \n", ccc, result[i*l+tmpI], len(result), i*l+ii)
	// 					result = append(result, append(result[i*l+tmpI], ccc))
	// 					tmpI++
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	ch := make(chan []int, l*l*l)
	for _, c := range nums {
		// fmt.Println("入队:", c)
		temp := make([]int, 0, l)
		temp = append(temp, c)
		ch <- temp
	}

	for temp := range ch {
		// fmt.Println("获取:", temp)
		tempI := -1
		for i, c := range nums {
			if c == temp[len(temp)-1] {
				tempI = i
				// fmt.Println("匹配到元素:", temp, " 容量:", cap(temp), " c:", c, " tempI:", tempI, " i:", i)
				result = append(result, temp)
			} else if tempI >= 0 && i > tempI {
				t := make([]int, len(temp))
				copy(t, temp)
				t = append(t, c)
				// fmt.Println("入队:", t, "基数:", nums[tempI], "补充:", c)
				ch <- t
				if len(t) == l {
					// fmt.Println("t:", len(t), " l:", l)
					close(ch)
				}
			}
		}

	}
	result = append(result, []int{})
	// fmt.Println(result)
	return result
}
