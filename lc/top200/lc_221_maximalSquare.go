package main

import "GoProject/leetcode/util"

/*
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
*/
func maximalSquare(matrix [][]byte) int {
	res := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1) // 存储以(i,j)为右下角的举行的最大边长
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + util.MinInt(util.MinInt(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1])
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}

	return res * res
}

func main() {

}
