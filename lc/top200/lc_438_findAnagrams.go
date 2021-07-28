package main

import "fmt"

/*
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指字母相同，但排列不同的字符串。
：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

*/
func findAnagrams(s string, p string) []int {
	var res []int

	need := make(map[byte]int, 128)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	l, valid := 0, 0
	for r := 0; r < len(s); r++ {
		need[s[r]]--
		if need[s[r]] >= 0 {
			valid++
		}

		for valid == len(p) {
			if r-l+1 == len(p) {
				res = append(res, l)
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
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}
