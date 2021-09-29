package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 环形的房子，首尾不能同时取
func rob_help(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	res, first, second := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			second = util.MaxInt(first, nums[i])
		} else {
			first, second = second, util.MaxInt(first+nums[i], second)
		}
		res = util.MaxInt(second, res)
	}
	return res
}

func rob213(nums []int) int {
	size := len(nums)
	if size < 1 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	res1 := rob_help(nums[:size-1])
	res2 := rob_help(nums[1:])
	return util.MaxInt(res1, res2)
}
func main() {
	nums := []int{1}
	fmt.Println(rob213(nums))
}
