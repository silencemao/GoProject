package main

import (
	"fmt"
	"GoProject/leetcode/util"
)
/*
最长上升子序列的长度，数组中的数值逐渐上升的子序列的长度，不要求连续
给定一个数组nums={10, 9, 2, 5, 4, 2, 9}，找出数组中最长上升子序列的长度

最长的长度，采用动态规划的方法，我们计算出以 i 位置的数为终点时的 最长上升子序列的长度，计算完所有位置时，就能找出最大的长度
i位置的长度如何获得，如果j（j<i）位置的数小于i位置的数，则在dp[j]的基础上加一 dp[j]+1，同时也要与i位置本身的值做对比，
即dp[i] = max(dp[j]+1, dp[i])
*/

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := -1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = util.MaxInt(dp[j] + 1, dp[i])
			}
		}
		res = util.MaxInt(dp[i], res)
	}
	return res
}

func main() {
	nums := []int{10, 9, 2, 3, 5, 7, 10, 18, 9, 20}
	fmt.Println(lengthOfLIS(nums))
}