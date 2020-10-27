package leetcode

import (
	"fmt"
)

//TODO
func findMedianSortedArrays4S1(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	target := l / 2
	even := false
	if l%2 == 0 {
		//even
		even = true
	} else {
		//odd
	}
	if l1 == 0 {
		if even {
			return (float64(nums2[target]) + float64(nums2[target-1])) / 2
		}
		return float64(nums2[target])
	} else if l2 == 0 {
		if even {
			return (float64(nums1[target]) + float64(nums1[target-1])) / 2
		}
		return float64(nums1[target])
	}
	for {
		t1 := len(nums1) / 2
		t2 := len(nums2) / 2
		fmt.Printf("--------%d--------\n", target)
		// fmt.Println("n1 mid location", t1, "n2 mid location", t2)
		fmt.Println("nums1", nums1, "nums2", nums2)
		fmt.Println("n1 mid value", nums1[t1], "n2 mid value", nums2[t2])
		if target <= 1 {
			break
		}
		if nums1[t1] >= nums2[t2] {
			target = target - t2
			nums1 = nums1[:t1+1]
			nums2 = nums2[t2:]
		} else {
			target = target - t1
			nums1 = nums1[t1:]
			nums2 = nums2[:t2+1]
		}
	}
	fmt.Println(even, "nums1", nums1, "nums2", nums2)
	if even {
		return (float64(nums1[0]) + float64(nums2[0])) / 2
	} else if len(nums1) == 1 {
		return float64(nums1[0])
	} else {
		return float64(nums2[0])
	}
}
