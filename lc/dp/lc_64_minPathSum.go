package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个包含非负整数的 mxn网格grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
*/
func minPathSum(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	tmp := make([][]int, m)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	tmp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				tmp[i][j] = grid[i][j] + util.MinInt(tmp[i-1][j], tmp[i][j-1])
			} else if i > 0 && j == 0 {
				tmp[i][j] = grid[i][j] + tmp[i-1][j]
			} else if i == 0 && j > 0 {
				tmp[i][j] = grid[i][j] + tmp[i][j-1]
			}
		}
	}
	res = tmp[m-1][n-1]
	return res
}

/*
滚动数组，降低空间复杂度
*/
func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = grid[i][j] + util.MinInt(dp[j-1], dp[j])
			}
		}
	}
	return dp[n-1]
}

func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(minPathSum(grid))
	fmt.Println(minPathSum1(grid))
}
