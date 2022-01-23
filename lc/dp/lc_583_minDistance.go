package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 当前位置不需要删除
			} else {
				dp[i][j] = util.MinInt(util.MinInt(dp[i-1][j], dp[i][j-1])+1, dp[i-1][j-1]+2) // 当前两个字符都删除 或只删除一个单词中的一个字符
			}
		}
	}
	return dp[m][n]
}

// 优化空间
func minDistance1(word1 string, word2 string) int {
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
				dp[i&1][j] = util.MinInt(dp[(i-1)&1][j-1]+2, util.MinInt(dp[(i-1)&1][j], dp[i&1][j-1])+1)
			}
		}
	}
	return dp[m&1][n]
}

func main() {
	word1 := "z"
	word2 := "mae"
	fmt.Println(minDistance(word1, word2))
	fmt.Println(minDistance1(word1, word2))
}
