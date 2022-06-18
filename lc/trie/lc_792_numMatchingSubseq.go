package main

import "fmt"

type node struct {
	word string
	ind  int
}

/*
给定字符串 s和字符串数组words, 返回words[i]中是s的子序列的单词个数。

字符串的 子序列 是从原始字符串中生成的新字符串，可以从中删去一些字符(可以是none)，而不改变其余字符的相对顺序。

例如， “ace” 是 “abcde” 的子序列。
输入: s = "abcde", words = ["a","bb","acd","ace"]
输出: 3
解释: 有三个是s 的子序列的单词: "a", "acd", "ace"。

1、words根据首字母进行分桶 {'a': [aw, arss], 'e':[ews, rfs]}，桶内每个单词由node组成，word表示该单词，ind表示该单词被遍历到第几个字母了
2、遍历s中每一个字母，找到对应的桶
3、遍历桶中每一个单词，若这个单词已经遍历到末尾，则表示该单词可以由s构成，
   若这个单词未遍历到末尾，则将这个单词根据字母放到新的桶中
*/
func numMatchingSubseq(s string, words []string) int {
	res := 0
	heads := make([][]*node, 26)
	for i := range heads {
		heads[i] = make([]*node, 0)
	}
	for _, w := range words {
		heads[w[0]-'a'] = append(heads[w[0]-'a'], &node{word: w, ind: 0})
	}

	for _, c := range s {
		oldHead := heads[c-'a']

		heads[c-'a'] = []*node{}
		for _, w := range oldHead {
			w.ind += 1
			if w.ind == len(w.word) {
				res += 1
				fmt.Println(w.word)
			} else {
				heads[w.word[w.ind]-'a'] = append(heads[w.word[w.ind]-'a'], w)
			}
		}
		oldHead = []*node{}
	}
	return res
}

func main() {
	s := "dsahjpjauf"
	words := []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}
	fmt.Println(numMatchingSubseq(s, words))
}
