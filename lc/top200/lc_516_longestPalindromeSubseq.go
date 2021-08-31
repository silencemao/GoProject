package main

import "GoProject/leetcode/util"

/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
*/
func longestPalindromeSubseq(s string) int {
	size := len(s)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = util.MaxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][size-1]
}
func main() {

}
