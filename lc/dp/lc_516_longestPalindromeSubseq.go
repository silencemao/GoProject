package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

*/
func longestPalindromeSubseq(s string) int {
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		dp[i][i] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := i; j < m; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = j - i + 1
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = util.MaxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][m-1]
}

func longestPalindromeSubseq1(s string) int {
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		dp[i][i] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := i + 1; j < m; j++ { // 数组初始化的时候，已经初始化dp[i][i]=1,所以j直接从i+1开始
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = util.MaxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][m-1]
}

func main() {
	s := "bbbbab"
	fmt.Println(longestPalindromeSubseq(s))
	fmt.Println(longestPalindromeSubseq1(s))
}
