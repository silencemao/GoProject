package main

import "fmt"

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	prev, cur, l1, l2 := 0, 0, 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		prev = cur
		if l2 < n && (l1 >= m || nums1[l1] > nums2[l2]) {
			cur = nums2[l2]
			l2 += 1
		} else {
			cur = nums1[l1]
			l1 += 1
		}
	}
	if (m+n)%2 == 1 {
		return float64(cur)
	}
	return float64(prev+cur) / 2
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
