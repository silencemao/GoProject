package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
同1143题
在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
现在，可以绘制一些连接两个数字 nums1[i]和 nums2[j]的直线，这些直线需要同时满足满足：
nums1[i] == nums2[j]
且绘制的直线不与任何其他连线（非水平线）相交。
请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

以这种方法绘制线条，并返回可以绘制的最大连线数。
*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
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
				dp[i][j] = util.MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
func main() {
	nums1, nums2 := []int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}
