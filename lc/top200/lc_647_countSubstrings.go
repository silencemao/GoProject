package main

import "fmt"

/*
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
注意遍历顺序
*/
func countSubstrings(s string) int {
	size := len(s)
	dp := make([][]bool, size+1)
	for i := range dp {
		dp[i] = make([]bool, size+1)
	}
	res := 0
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if s[i-1] == s[j-1] {
				if j-i <= 1 {
					dp[i][j] = true
					res++
				} else {
					if dp[i+1][j-1] {
						dp[i][j] = true
						res++
					} else {
						dp[i][j] = false
					}
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return res
}
func main() {
	s := "abcbaadefhf"
	fmt.Println(countSubstrings(s))
}
