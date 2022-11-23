package main

import "fmt"

type node struct {
	word string
	ind  int
}

/*
匹配子序列的单词数
给定字符串 s和字符串数组words, 返回words[i]中是s的子序列的单词个数。
字符串的 子序列 是从原始字符串中生成的新字符串，可以从中删去一些字符(可以是none)，而不改变其余字符的相对顺序。
例如， “ace” 是 “abcde” 的子序列。

输入: s = "abcde", words = ["a","bb","acd","ace"]
输出: 3
解释: 有三个是s 的子序列的单词: "a", "acd", "ace"。

*/
func numMatchingSubseq(s string, words []string) int {
	ans := 0
	heads := make([][]*node, 26)
	for i := range heads {
		heads[i] = make([]*node, 0)
	}
	for _, w := range words {
		heads[w[0]-'a'] = append(heads[w[0]-'a'], &node{word: w, ind: 0})
	}
	for _, c := range s {
		oldHead := heads[c-'a']
		heads[c-'a'] = nil

		for _, w := range oldHead {
			w.ind += 1
			if w.ind == len(w.word) {
				ans += 1
			} else {
				heads[w.word[w.ind]-'a'] = append(heads[w.word[w.ind]-'a'], &node{word: w.word, ind: w.ind})
			}
		}
	}
	return ans
}

func main() {
	s := "abcde"
	words := []string{"a", "bb", "acd", "ace"}
	fmt.Println(numMatchingSubseq(s, words))
}
