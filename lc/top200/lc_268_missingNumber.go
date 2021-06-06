package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 [0, n]中n个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?
https://leetcode-cn.com/problems/missing-number
*/
func missingNumber(nums []int) int {
	sort.Ints(nums)
	pos := 0
	for ; pos < len(nums); pos++ {
		if nums[pos] != pos {
			return pos
		}
	}
	return pos
}

func missingNumber1(nums []int) int {
	m := make([]int, len(nums)+1)
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = nums[i]
	}
	i := 0
	for ; i < len(m); i++ {
		if m[i] != i {
			return i
		}
	}
	return i
}

func missingNumber2(nums []int) int {
	res := 0
	res ^= len(nums)
	for i, num := range nums {
		res = res ^ i ^ num
	}
	return res
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber2(nums))
}
