package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

分两次求解，第一次包含第一家(i=0)，不包含最后一家(i=size-1)
           第二次不包含第一家，包含最后一家
两次求解结果取最大
*/
func help213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second, res := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			second = util.MaxInt(first, nums[i])
		} else {
			first, second = second, util.MaxInt(first+nums[i], second)
		}
		res = util.MaxInt(res, second)
	}
	return res
}

func rob213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res1 := help213(nums[1:])
	res2 := help213(nums[:len(nums)-1])
	if res1 > res2 {
		return res1
	}
	return res2
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(rob213(nums))
}
