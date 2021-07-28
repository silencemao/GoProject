package main

import (
	"fmt"
)

/*
给定一个字符串s和一些 长度相同 的单词words 。找出 s 中恰好可以由words 中所有单词串联形成的子串的起始位置。
注意子串要与words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑words中单词串联的顺序。
https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words

s = "barfoothefoobarman", words = ["foo","bar"]
[0,9]
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。

注意 若words=["for","bao"] s中不存在子串可以由words组成，返回[]
words中每个单词的长度相等

*/
func findSubstring(s string, words []string) []int {
	var res []int
	l, valid, size, wordNum := 0, 0, len(words[0]), len(words)
	need := make(map[byte]int, 128)    // 记录words中每个字母出现的次数
	subWord := make(map[string]int, 0) // 记录words中每个单词出现的次数
	for _, word := range words {
		for _, c := range word {
			need[byte(c)]++
		}
		subWord[word] += 1
	}

	for r := 0; r < len(s); r++ {
		need[s[r]]--
		if need[s[r]] >= 0 {
			valid++
		}

		for valid == wordNum*size {
			if r-l+1 == wordNum*size {
				tmp := make(map[string]int, 0) // 记录s中 l-r范围内 每个单词出现的次数
				for m := l; m <= r; m += size {
					tmp[s[m:m+size]] += 1
				}
				isOK := true
				for key := range tmp { // l-r中每个单词出现的次数是否和subword中一致
					if subWord[key] != tmp[key] {
						isOK = false
					}
				}
				if isOK {
					res = append(res, l)
				}
			}
			need[s[l]]++
			if need[s[l]] > 0 {
				valid--
			}
			l++
		}
	}
	return res
}
func main() {
	s := "barfoothefoobarman"
	words := []string{"bar", "foo"}
	fmt.Println(findSubstring(s, words))
}
