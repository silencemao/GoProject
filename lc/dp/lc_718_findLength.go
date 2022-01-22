package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给两个整数数组A和B，返回两个数组中公共的、长度最长的子数组的长度。

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
长度最长的公共子数组是 [3, 2, 1] 。
*/
func findLength(nums1 []int, nums2 []int) int {
	res := 0
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			res = util.MaxInt(dp[i][j], res)
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(nums1, nums2))
}
