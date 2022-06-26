package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为11（即，2+3+5+1= 11）。

*/
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		n := len(triangle[i])
		dp[i] = make([]int, n)
		if i == 0 {
			dp[0][0] = triangle[0][0]
			continue
		}
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			} else if j == n-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = util.MinInt(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	res := 1<<31 - 1
	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		if dp[m-1][j] < res {
			res = dp[m-1][j]
		}
	}
	return res
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
