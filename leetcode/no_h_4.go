package leetcode

//TODO log(m+n)
func findMedianSortedArrays4S1(nums1 []int, nums2 []int) float64 {
	return 0
}

//DONE
func findMedianSortedArrays4S2(nums1 []int, nums2 []int) float64 {
	odd := false
	l1, l2 := len(nums1), len(nums2)
	if (l1+l2)&1 == 1 {
		odd = true
	}
	step := (l1 + l2 - 1) / 2
	for step != -1 {
		v1, v2 := -1, -1
		if len(nums1) > 0 {
			v1 = nums1[len(nums1)-1]
		}
		if len(nums2) > 0 {
			v2 = nums2[len(nums2)-1]
		}
		if step == 0 {
			vv := -1
			if odd {
				if v1 > v2 {
					return float64(v1)
				}
				return float64(v2)
			}
			if v1 >= v2 {
				nums1 = nums1[:len(nums1)-1]
				if len(nums1) > 0 {
					vv = nums1[len(nums1)-1]
				}
				if vv > v2 {
					return (float64(v1) + float64(vv)) / 2
				}
				return (float64(v1) + float64(v2)) / 2
			}
			nums2 = nums2[:len(nums2)-1]
			if len(nums2) > 0 {
				vv = nums2[len(nums2)-1]
			}
			if vv > v1 {
				return (float64(vv) + float64(v2)) / 2
			}
			return (float64(v1) + float64(v2)) / 2
		}
		if v1 >= v2 {
			nums1 = nums1[:len(nums1)-1]
		} else if v1 < v2 {
			nums2 = nums2[:len(nums2)-1]
		}
		step--
	}
	return -1
}
