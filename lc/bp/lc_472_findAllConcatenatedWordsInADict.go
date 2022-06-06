package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*

给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。

连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。

输入：words = ["cat","dog","catdog"]
输出：["catdog"]
*/
/*
超时
*/
func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var res []string
	minL := len(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) < 2*minL {
			continue
		}
		word := words[i]
		dp := make([]bool, len(word)+1)
		dp[0] = true
		for m := 0; m <= len(word); m++ {
			for n := 0; n < i && len(words[n]) < len(word); n++ { // 备选单词索引
				if m >= len(words[n]) && word[m-len(words[n]):m] == words[n] {
					dp[m] = dp[m] || dp[m-len(words[n])]
				}
			}
		}
		if dp[len(dp)-1] {
			res = append(res, word)
		}
	}
	return res
}

/*
https://leetcode.cn/problems/concatenated-words/solution/tu-jie-kun-nan-ti-mu-kun-nan-ti-qi-shi-h-4p6z/
1、先对所有单词生成字典
2、每个单词判断，是否可以有两个以及以上单词组成
3、
*/
func findAllConcatenatedWordsInADict1(words []string) []string {
	set := make(map[string]bool, 0)
	for _, word := range words {
		set[word] = true
	}
	var res []string

	for _, word := range words {
		if help472(word, set) {
			res = append(res, word)
		}
	}
	return res
}

func help472(word string, set map[string]bool) bool {
	n := len(word)
	dp := make([]int, n+1) // dp[i]表示i以及i之前的位置可以由几个单词组成，dp[i]=-1表示无效，不能拆分成多个单词
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		if dp[i] == -1 { // dp[i] == -1 表示i位置前面的单词不能由set中的单词组成，所以i以及前面位置不可以拆开，因此不能从i位置开始向后找单词拆分
			continue
		}
		for j := i + 1; j <= n; j++ {
			cur := word[i:j]
			if set[cur] {
				dp[j] = util.MaxInt(dp[j], dp[i]+1)
			}
		}
		if dp[n] > 1 {
			return true
		}
	}
	return false
}

func main() {
	words := []string{"a", "b", "ab", "abc"}
	fmt.Println(findAllConcatenatedWordsInADict(words))
	fmt.Println(findAllConcatenatedWordsInADict1(words))
}
