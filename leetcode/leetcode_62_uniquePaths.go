package main

import "fmt"

/*
给定一个二维矩阵，起始点在左上角，每次只能向右或者向下走，问走到右下角右多少种方式

一般这种最多右多少种方式，最大值的问题都可以考虑dp

走到矩阵中任意一个点有几种方式取决于走到这个点的上面一个点和左边一个点有几种方式
即 dp[i][j] = dp[i-1][j] + dp[i][j-1]
第一行和第一列单独考虑，因为他们只有一种方式
*/

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	if m==1 || n==1 {
		return 1
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i==0 || j==0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	m, n := 7, 3
	fmt.Println(uniquePaths(m, n))
}
