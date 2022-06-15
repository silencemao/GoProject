package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
	"sort"
)

/*给出一个字符串数组 words 组成的一本英语词典。返回words 中最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。

若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。



示例 1：

输入：words = ["w","wo","wor","worl", "world"]
输出："world"
解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。

*/
func longestWord(words []string) string {
	trie := _struct.Constructor()
	for _, w := range words {
		trie.Insert(w)
	}
	res := ""
	for _, w := range words {
		if trie.Search(w) {
			if len(w) > len(res) {
				res = w
			} else if len(w) == len(res) && w < res {
				res = w
			}
		}
	}
	return res
}

func main() {
	words := []string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		} else if len(words[i]) > len(words[j]) {
			return false
		} else {
			if words[i] < words[j] {
				return true
			} else {
				return false
			}
		}
	})
	fmt.Println(words)
	fmt.Println(longestWord(words))
}
