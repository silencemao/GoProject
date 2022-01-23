package main

import (
	"fmt"
)

/*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
回文字符串 是正着读和倒过来读一样的字符串。
子字符串 是字符串中的由连续字符组成的一个序列。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
*/
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(i, j, s) {
				res += 1
			}
		}
	}
	return res
}

func isPalindrome(l, r int, s string) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func countSubstrings1(s string) int {
	res := 0
	m := len(s)
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}
	for i := m - 1; i >= 0; i-- {
		for j := i; j < m; j++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i <= 1 { // 子串长度为1(a) 或 为2(aa)
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] {
				res += 1
			}
		}
	}

	return res
}

func main() {
	s := "aaa"
	fmt.Println(countSubstrings1(s))
}
