package main

import (
	"fmt"
)

/*
单词拆分 输出所有拆分后的结果

递归求解
*/
func wordBreakII(s string, wordDict []string) []string {
	m := make(map[string][]string)
	return helper140(m, s, wordDict)
}
func helper140(m map[string][]string, s string, wordDict []string) []string {
	if tmp, ok := m[s]; ok {
		return tmp
	}
	res := []string{}
	if len(s) == 0 {
		return []string{""}
	}

	for _, word := range wordDict {
		if len(word) > len(s) || s[:len(word)] != word {
			continue
		}
		words := helper140(m, s[len(word):], wordDict)
		for _, w := range words {
			if len(w) > 0 {
				res = append(res, word+" "+w)
			} else {
				res = append(res, word+w)
			}
		}
	}
	m[s] = res
	return m[s]
}

func main() {
	s := "aaaaaaa"
	wordDict := []string{"aaaa", "aa", "a"}
	res := wordBreakII(s, wordDict)
	for _, r := range res {
		fmt.Println(r)
	}
	fmt.Println(len(res))
}
