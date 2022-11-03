package main

import "fmt"

/*
最大重复子字符串
给你一个字符串sequence，如果字符串 word连续重复k次形成的字符串是sequence的一个子字符串，
那么单词word 的 重复值为 k 。单词 word的 最大重复值是单词word在sequence中最大的重复值。
如果word不是sequence的子串，那么重复值k为 0 。

给你一个字符串 sequence和 word，请你返回 最大重复值k 。

输入：sequence = "ababc", word = "ab"
输出：2
解释："abab" 是 "ababc" 的子字符串。

*/
func maxRepeating(sequence string, word string) int {
	m, n := len(sequence), len(word)
	dp := make([]int, m)
	res := 0
	for i := n - 1; i < m; i++ {
		if sequence[i-n+1:i+1] == word {
			if i == n-1 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-n] + 1
			}
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}

func main() {
	sequence := "ababc"
	word := "ba"
	fmt.Println(maxRepeating(sequence, word))
}
