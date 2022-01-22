package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，
那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

应该叫连续递增子数组，和第300题不一样，本题要求是连续递增子数组，num[i] <= num[i-1]，则num[i]不能和i之前的数字构成连续子数组了，所以dp[i]=1
表示当前位置结尾的连续递增子数组的长度为1

*/

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		res = util.MaxInt(dp[i], res)
	}
	return res
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}
