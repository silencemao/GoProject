package main

import "fmt"

// 给定矩阵的大小，起点在左上角 终点在右下角
// 从起点到终点共多少种走法
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 对上面做法的优化
func uniquePaths1(m, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1] // dp[j] = dp[j](上边) + dp[j-1](左边)
		}
	}
	return dp[n-1]
}

func main() {
	m, n := 3, 3
	fmt.Println(uniquePaths(m, n))
	fmt.Println(uniquePaths1(m, n))
}
