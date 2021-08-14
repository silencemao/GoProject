package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，
那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
注意：是连续递增子序列，要求 l-r之间的索引不间断（相当于子数组）

https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence

*/

// 这种写法不对，不是连续子数组 nums = [1, 3, 5, 4, 7] 会返回4(1 3 5 7) 5和7之间有4间隔着，不满足题意
func findLengthOfLCIS(nums []int) int {
	n, res := len(nums), 0
	if n < 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = util.MaxInt(dp[j]+1, dp[i])
			} else if nums[i] == nums[j] {
				dp[i] = util.MaxInt(dp[j], dp[i])
			} else {
				// 继续向前走
			}
		}
		res = util.MaxInt(res, dp[i])
	}
	fmt.Println(dp)
	return res
}

// dp 数组 dp[i] = max(dp[i-1]+1, dp[i])
// 注意当nums[i] <= nums[i-1]时，dp[i]=1
func findLengthOfLCIS2(nums []int) int {
	n, res := len(nums), 0
	if n < 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		res = util.MaxInt(dp[i], res)
	}
	fmt.Println(dp)
	return res
}

func findLengthOfLCIS3(nums []int) int {
	n, res := len(nums), 0
	if n < 2 {
		return n
	}
	cur := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cur += 1
		} else {
			cur = 1
		}
		res = util.MaxInt(res, cur)
	}
	return res
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS3(nums))
}
