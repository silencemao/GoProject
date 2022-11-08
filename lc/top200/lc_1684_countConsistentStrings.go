package main

import "fmt"

/*
统计一致字符串的数目
给你一个由不同字符组成的字符串allowed和一个字符串数组words。如果一个字符串的每一个字符都在 allowed中，就称这个字符串是 一致字符串 。
请你返回words数组中一致字符串 的数目。
输入：allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
输出：2
解释：字符串 "aaab" 和 "baa" 都是一致字符串，因为它们只包含字符 'a' 和 'b' 。
*/
func countConsistentStrings(allowed string, words []string) int {
	res := 0
	set := make(map[int]bool)
	for _, b := range allowed {
		set[int(b-'a')] = true
	}
	for _, word := range words {
		cnt := 0
		for _, b := range word {
			if _, ok := set[int(b-'a')]; ok {
				cnt += 1
			}
		}
		if cnt == len(word) {
			res += 1
		}
	}
	return res
}
func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	fmt.Println(countConsistentStrings(allowed, words))
}
