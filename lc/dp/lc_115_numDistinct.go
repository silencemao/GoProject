package main

import "fmt"

/*
给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。

字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（
例如，"ACE"是"ABCDE"的一个子序列，而"AEC"不是）

题目数据保证答案符合 32 位带符号整数范围

dp[i][j]表示以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]。
dp[i][0]表示以i-1为结尾的s子序列中出现以j-1为结尾的(空字符串)个数，以i-1为结尾的字符串s可以删除元素，得到空字符串 所以dp[i][0]=1
dp[0][j]表示s为空字符串时，可以组成以j-1为结尾的t字符串的个数 dp[0][j]=0
*/

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j] // s=bagg t=bag，s中保留字符s[i-1] + s中不保留字符s[i-1]可以组成字符串t的个数
			} else {
				dp[i][j] = dp[i-1][j] // 删除s[i-1]
			}
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return dp[m][n]
}

func main() {
	s := "bagg"
	t := "bag"
	fmt.Println(numDistinct(s, t))
}
