package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)

/*
给定二维三角数组，计算从第一行到最后一行路径之和的最小值
*/

/*
动态规划 从顶至下
*/
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = util.MinInt(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	res := 1<<31 - 1
	for i := 0; i < len(dp[m-1]); i++ {
		if dp[m-1][i] < res {
			res = dp[m-1][i]
		}
	}
	return res
}

/*
动态规划，从下至顶

bottom to up相比于 up to bottom的优点在于可以使用一维dp

在计算第layer 行的时候，使用一维dp记录layer+1行的路径之和，(从最后一行至layer+1行的路径之和)，
然后再将第layer行的i个数与[layer+1][i]和[layer+1]中的较小值相加，而layer+1行的结果就存在一维数组中，
我们可以直接在一维数组中更新
*/
func minimumTotal1(triangle [][]int) int {
	var minPath []int
	if len(triangle) == 0 || len(triangle[len(triangle)-1]) == 0 {
		return 1<<31 - 1
	}
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		minPath = append(minPath, triangle[len(triangle)-1][i])
	}
	for layer := len(triangle) - 2; layer >= 0; layer-- {
		for i := 0; i <= layer; i++ {
			minPath[i] = util.MinInt(minPath[i], minPath[i+1]) + triangle[layer][i]
		}
	}
	return minPath[0]
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
	fmt.Println(minimumTotal1(triangle))
}
