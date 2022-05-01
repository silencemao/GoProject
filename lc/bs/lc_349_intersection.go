package main

import (
	"fmt"
	"sort"
)

/*
给定两个数组nums1和nums2 ，返回 它们的交集。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
*/
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	pre, ind1, ind2, l1, l2 := -1234, 0, 0, len(nums1), len(nums2)
	var res []int
	for ind1 < l1 && ind2 < l2 {
		if nums1[ind1] == nums2[ind2] {
			if pre != nums2[ind2] {
				res = append(res, nums1[ind1])
				pre = nums1[ind1]
			}
			ind1++
			ind2++
		} else if nums1[ind1] < nums2[ind2] {
			ind1++
		} else {
			ind2++
		}
	}
	return res
}

func intersection1(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	if m > n {
		return intersection(nums2, nums1)
	}
	set := make(map[int]int, 0)
	for _, num := range nums1 {
		if _, ok := set[num]; !ok {
			set[num] = 1
		}
	}
	res := []int{}
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			res = append(res, num)
			delete(set, num)
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
