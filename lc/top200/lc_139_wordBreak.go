package main

import "fmt"

/*
超时
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。

遍历每个单词，在s中逐个位置查找，是否存在一个子串与word相等，若相等，则继续判断子串左边 与 右边两个子串是否也可以有worddict中的单词组成，
若左右都能由wordDict中的单词组成，则表示可以正确拆分
*/
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || len(wordDict) == 0 {
		return true
	}
	for i := 0; i < len(wordDict); i++ {
		for j := 0; j+len(wordDict[i]) <= len(s); j++ {
			if s[j:j+len(wordDict[i])] == wordDict[i] {
				if wordBreak(s[0:j], wordDict) && wordBreak(s[j+len(wordDict[i]):], wordDict) {
					return true
				}
			}
		}
	}
	return false
}

/*
转变成完全背包问题，（单词可以重复使用）
以s中每一个位置为结尾，看是否存在一个word 与 s中以i为结尾的子串相等，
若相等，我们看以i结尾的子串左侧的子串是否仍然能由 wordDict中的word组成，
若能，表示从0-i长度的子串，均可有wordDict中的单词组成
转移方程为 dp[i] = dp[i-l] || dp[i]
  01 234
  ca rs
     rs
以4为结尾的子串rs可以由wordDict中的单词组成，看以[4-2]=2位置结尾的子串也可以由wordDict中的子串组成，
说明从0-4之间的子串可以由wordDict中的单词组成
*/
func wordBreak1(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i <= len(s); i++ {
		for _, word := range wordDict {
			l := len(word)
			if i-l >= 0 && s[i-l:i] == word {
				dp[i] = dp[i] || dp[i-l]
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "cars"
	wordDict := []string{"car", "ca", "rs"}
	fmt.Println(wordBreak(s, wordDict))
	fmt.Println(wordBreak1(s, wordDict))
}
