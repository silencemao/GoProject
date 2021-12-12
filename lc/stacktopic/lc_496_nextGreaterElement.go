package main

import "fmt"

/*
nums1中数字x的 下一个更大元素 是指x在nums2 中对应位置 右侧 的 第一个 比x大的元素。
给你两个 没有重复元素 的数组nums1 和nums2 ，下标从 0 开始计数，其中nums1是nums2的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。
如果不存在下一个更大元素，那么本次查询的答案是 -1 。
返回一个长度为nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

单调栈
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res []int
	for i, v := range nums1 {
		isOk := false
		for _, v2 := range nums2 {
			if isOk && v2 > v {
				res = append(res, v2)
				break
			}
			if v == v2 {
				isOk = true
			}
		}
		if !isOk || len(res) < i+1 {
			res = append(res, -1)
		}
	}
	return res
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	set := make(map[int]int, 0)
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] { // 此处把所有小于nums[i]的索引都删除
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			set[nums2[top]] = nums2[i]
		}
		stack = append(stack, i)
	}

	for i, v := range nums1 {
		if v2, ok := set[v]; ok {
			res[i] = v2
		} else {
			res[i] = -1
		}
	}

	return res
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	set := make(map[int]int, 0)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			set[nums2[i]] = nums2[stack[len(stack)-1]]
		} else {
			set[nums2[i]] = -1
		}
		stack = append(stack, i)
	}
	for i, v := range nums1 {
		if v2, ok := set[v]; ok {
			res[i] = v2
		} else {
			res[i] = -1
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 3, 5, 2, 4}
	nums2 := []int{6, 5, 4, 3, 2, 1, 7}
	fmt.Println(nextGreaterElement(nums1, nums2))
	fmt.Println(nextGreaterElement1(nums1, nums2))
	fmt.Println(nextGreaterElement2(nums1, nums2))
}
