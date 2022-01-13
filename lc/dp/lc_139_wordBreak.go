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
