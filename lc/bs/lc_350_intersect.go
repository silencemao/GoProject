package main

import (
	"fmt"
	"sort"
)

/*
给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。
返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。
可以不考虑输出结果的顺序。
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
返回[2,2]
第349题的升级版
*/
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	_, ind1, ind2, l1, l2 := -1234, 0, 0, len(nums1), len(nums2)
	var res []int
	for ind1 < l1 && ind2 < l2 {
		if nums1[ind1] == nums2[ind2] {
			//if pre != nums2[ind2] {
			res = append(res, nums1[ind1])
			//pre = nums1[ind1]
			//}
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

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}
