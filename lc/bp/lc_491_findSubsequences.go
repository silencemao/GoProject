package main

import (
	"fmt"
)

/*
给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。

示例 1：

输入：nums = [4,6,7,7]
输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

*/

func help491(res *[][]int, nums, tmp []int, ind int) {
	if len(tmp) >= 2 {
		*res = append(*res, append([]int{}, tmp...))
	}
	used := make([]int, 201)
	for i := ind; i < len(nums); i++ {
		if used[nums[i]+100] == 1 || (len(tmp) > 0 && nums[i] < tmp[len(tmp)-1]) {
			continue
		}

		used[nums[i]+100] = 1
		tmp = append(tmp, nums[i])
		help491(res, nums, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func findSubsequences(nums []int) [][]int {
	var res [][]int
	help491(&res, nums, []int{}, 0)
	return res
}

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))
}
