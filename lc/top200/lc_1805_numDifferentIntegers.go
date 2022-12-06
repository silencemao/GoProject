package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
字符串中不同整数的数目

给你一个字符串 word ，该字符串由数字和小写英文字母组成。

请你用空格替换每个不是数字的字符。例如，"a123bc34d8ef34" 将会变成 " 123 34 8 34" 。注意，剩下的这些整数为（相邻彼此至少有一个空格隔开）："123"、"34"、"8" 和 "34" 。
返回对 word 完成替换后形成的 不同 整数的数目。
只有当两个整数的 不含前导零 的十进制表示不同， 才认为这两个整数也不同。

输入：word = "a123bc34d8ef34"
输出：3
解释：不同的整数有 "123"、"34" 和 "8" 。注意，"34" 只计数一次。
*/
func numDifferentIntegers(word string) int {
	res := 0
	ss := []byte(word)
	for i := 0; i < len(ss); i++ {
		if ss[i] < '0' || ss[i] > '9' {
			ss[i] = ' '
		}
	}
	word = string(ss)
	set := map[int]bool{}
	for _, w := range strings.Split(word, " ") {
		num, err := strconv.Atoi(w)
		if err != nil {
			continue
		}
		if set[num] {
			continue
		} else {
			set[num] = true
			res += 1
		}
	}
	return res
}

func numDifferentIntegers1(word string) int {
	s, n := map[string]struct{}{}, len(word)

	for i := 0; i < n; i++ {
		if word[i] >= '0' && word[i] <= '9' {
			for i < n && word[i] == '0' {
				i += 1
			}
			j := i
			for j < n && word[j] >= '0' && word[j] <= '9' {
				j += 1
			}
			s[word[i:j]] = struct{}{}
			i = j
		}
	}

	return len(s)
}

func main() {
	word := "0a0"
	fmt.Println(numDifferentIntegers(word))
	fmt.Println(numDifferentIntegers1(word))

}
