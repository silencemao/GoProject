package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
	"strings"
)

/*
在英语中，我们有一个叫做词根(root) 的概念，可以词根后面添加其他一些词组成另一个较长的单词——我们称这个词为继承词(successor)。
例如，词根an，跟随着单词other(其他)，可以形成新的单词another(另一个)。
现在，给定一个由许多词根组成的词典 dictionary 和一个用空格分隔单词形成的句子 sentence。你需要将句子中的所有继承词用词根替换掉。
如果继承词有许多可以形成它的词根，则用最短的词根替换它。

你需要输出替换之后的句子。

输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"

将sentence中的单词替换为以dictionary中开头的单词
*/
func replaceWords(dictionary []string, sentence string) string {
	set := make(map[string]bool, 0)
	for _, w := range dictionary {
		set[w] = true
	}
	words := strings.Split(sentence, " ")
	for i, w := range words {
		for j := range w {
			if set[w[:j+1]] {
				words[i] = w[:j+1]
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func replaceWords1(dictionary []string, sentence string) string {
	trie := _struct.Constructor()
	for _, w := range dictionary {
		trie.Insert(w)
	}
	words := strings.Split(sentence, " ")
	for i, w := range words {

		for j := 0; j < len(w); j++ {
			if trie.Search(w[:j+1]) {
				words[i] = w[:j+1]
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
	fmt.Println(replaceWords1(dictionary, sentence))
}
