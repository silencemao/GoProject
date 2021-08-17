package main

import (
	"fmt"
)

/*
给定字符串 s t，判断s是否为t的一个子序列，t中的字符可以删除也可以不删除
和1143题类似，相当于找最长公共子序列，且最长公共子序列的长度与s的长度一致
注意处理边界值
*/
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	if m == 0 {
		return true
	}
	if m > n {
		return false
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1] // 相当于删除t中的字符
			}
			if dp[i][j] == m {
				return true
			}
		}
	}
	return false
}
func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
