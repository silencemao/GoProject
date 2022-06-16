package main

import (
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

type Trie struct {
	IsEnd bool
	Next  []*Trie
}

func Constructor() Trie {
	trie := Trie{}
	trie.IsEnd = false
	trie.Next = make([]*Trie, 26)
	return trie
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		ch := c - 'a'
		if node.Next[ch] == nil {
			t := Constructor()
			node.Next[ch] = &t
		}
		node = node.Next[ch]
	}
	node.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		ch := c - 'a'
		node = node.Next[ch]
		if node == nil || !node.IsEnd { // 控制每个前缀是由列表中其它单词加一个字母得到
			return false
		}
	}
	return node.IsEnd
}

func longestWord(words []string) string {
	trie := Constructor()
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
