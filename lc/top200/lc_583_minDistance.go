package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

//https://mp.weixin.qq.com/s/a8BerpqSf76DCqkPDJrpYg
func minDistance583(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = util.MinInt(util.MinInt(dp[i-1][j], dp[i][j-1])+1, dp[i-1][j-1]+2)
			}
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return dp[m][n]
}

/*
https://leetcode-cn.com/problems/delete-operation-for-two-strings/solution/javajie-fa-er-wei-dong-tai-gui-hua-gun-d-or8d/
滚动数组
*/
func minDistance583_2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		dp[i&1][0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i&1][j] = dp[(i-1)&1][j-1]
			} else {
				dp[i&1][j] = util.MinInt(dp[(i-1)&1][j-1]+2, util.MinInt(dp[i&1][j-1], dp[(i-1)&1][j])+1)
			}
		}
	}
	return dp[m&1][n]
}

func main() {
	word1, word2 := "sea", "eat"
	fmt.Println(minDistance583(word1, word2))
	fmt.Println(minDistance583_2(word1, word2))
}
