package leetcode

//DONE
func intersection349(nums1 []int, nums2 []int) []int {
	dic := make(map[int]bool)
	for _, k := range nums1 {
		if _, ok := dic[k]; !ok {
			dic[k] = true
		}
	}
	result := []int{}
	for _, k := range nums2 {
		if _, ok := dic[k]; ok {
			dic[k] = false
		}
	}
	for k, v := range dic {
		if !v {
			result = append(result, k)
		}
	}
	return result
}
