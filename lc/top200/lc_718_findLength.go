package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。
*/
func findLength(nums1 []int, nums2 []int) int {
	res := 0
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0 // 若不相等，从i j就要重新开始算起，不能和前面相等的部分连起来
			}
			res = util.MaxInt(dp[i][j], res)
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return res
}

/*
状态压缩 改为滚动数组 一维
*/
func findLength2(nums1 []int, nums2 []int) int {
	res := 0
	m, n := len(nums1), len(nums2)
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0 // 若不相等，从i j就要重新开始算起，不能和前面相等的部分连起来
			}
			res = util.MaxInt(dp[j], res)
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(nums1, nums2))
}
