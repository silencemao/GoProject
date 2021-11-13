package main

import "fmt"

/*
返回两个数组的交集，且返回结果中不存在重复元素
[1, 2, 2, 1] [2, 2]
返回 [2]
*/
func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	set := make(map[int]int, 0)
	for _, num := range nums1 {
		if _, ok := set[num]; !ok {
			set[num] = 1
		}
	}
	for _, num := range nums2 {
		if cnt, ok := set[num]; ok && cnt > 0 {
			set[num] -= 1
			res = append(res, num)
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
