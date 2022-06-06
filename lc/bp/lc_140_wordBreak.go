package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串 s 和一个字符串字典wordDict，在字符串s中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
注意：词典中的同一个单词可能在分段中被重复使用多次。

示例 1：
输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
输出:["cats and dog","cat sand dog"]

单词拆分，将字符串拆分成多个单词，每个单词都在wordDict中出现，返回所有拆分方案
*/
func help140(res *[]string, tmp []string, set map[string]bool, s string, ind int) {
	if ind >= len(s) {
		*res = append(*res, strings.Join(tmp, " "))
	}
	for i := ind; i < len(s); i++ {
		if !set[s[ind:i+1]] {
			continue
		}
		tmp = append(tmp, s[ind:i+1])
		help140(res, tmp, set, s, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func wordBreak(s string, wordDict []string) []string {
	var res []string
	set := make(map[string]bool, 0)
	for _, word := range wordDict {
		set[word] = true
	}
	help140(&res, []string{}, set, s, 0)
	return res
}

func main() {

	s := "pineapplepenapple"
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	fmt.Println(wordBreak(s, wordDict))
}
