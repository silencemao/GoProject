package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
https://leetcode-cn.com/problems/longest-increasing-subsequence
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = util.MaxInt(dp[i], dp[j]+1)
			}
		}
		res = util.MaxInt(dp[i], res)
	}
	return res
}

//https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-e/
// 动态规划+二分的思路没看懂
func main() {
	nums := []int{0, 3, 1, 6, 2, 7}
	fmt.Println(lengthOfLIS(nums))
}
