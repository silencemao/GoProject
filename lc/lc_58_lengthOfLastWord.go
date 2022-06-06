package main

import (
	"fmt"
	"strings"
)

/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为5。

*/
func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) > 0 {
			//fmt.Println(words[i])
			return len(words[i])
		}
	}

	return 0
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
