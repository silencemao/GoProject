package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for _, word := range wordDict {
		for j := len(word); j < len(s)+1; j++ {
			if s[j-len(word):j] == word {
				dp[j] = dp[j] || dp[j-len(word)]
			}
		}
	}
	return dp[len(dp)-1]
}

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
"aaaaaaa"
["aaaa","aaa"]
*/
func wordBreak1(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for j := 0; j <= len(s); j++ {
		for _, word := range wordDict {
			if j >= len(word) && word == s[j-len(word):j] {
				dp[j] = dp[j] || dp[j-len(word)]
			}
		}
	}
	return dp[len(dp)-1]
}

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
	fmt.Println(wordBreak1(s, wordDict))
}
