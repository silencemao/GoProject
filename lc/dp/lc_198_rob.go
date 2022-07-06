package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
打家劫舍 相邻的位置不能同时打劫
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = util.MaxInt(nums[0], nums[1])
	res := dp[1]
	for i := 2; i < len(nums); i++ {
		for j := i - 2; j >= 0; j-- {
			dp[i] = util.MaxInt(dp[j]+nums[i], dp[i])
		}
		res = util.MaxInt(res, dp[i])
	}
	fmt.Println(dp)
	return res
}

func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res, first, second := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			second = util.MaxInt(first, nums[i])
		} else {
			first, second = second, util.MaxInt(first+nums[i], second)
		}
		res = util.MaxInt(res, second)
		// fmt.Println(res)
	}
	return res
}

func rob2(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	res, dp[0], dp[1] = util.MaxInt(nums[0], nums[0]), nums[0], util.MaxInt(nums[0], nums[1])
	for i := 2; i < len(dp); i++ {
		dp[i] = util.MaxInt(dp[i-1], dp[i-2]+nums[i])
		res = util.MaxInt(dp[i], res)
	}
	fmt.Println(dp)
	return res
}

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
	fmt.Println(rob1(nums))
	fmt.Println(rob2(nums))
}
