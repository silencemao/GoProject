package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。
*/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = util.MaxInt(dp[i-1]+nums[i], nums[i])
		res = util.MaxInt(dp[i], res)
	}
	fmt.Println(dp)
	return res
}

func maxSubArray1(nums []int) int {
	curMax, res := 0, 0
	for i := 0; i < len(nums); i++ {
		curMax = util.MaxInt(curMax+nums[i], nums[i])
		res = util.MaxInt(res, curMax)
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
