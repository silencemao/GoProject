package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串s ，找出s中最长的子字符串，子串中每个字符出现的次数不小于k

递归，找到子串中不满足长度的字符，以该字符为分隔符，分成的子串分别去判断，取最长的子串长度即可
*/
func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	res := 0
	set := make(map[int32]int, 0)
	for _, c := range s {
		set[c-'a'] += 1
	}
	for c := range set {
		if set[c] < k {
			for _, subStr := range strings.Split(s, string(c+'a')) {
				if len(subStr) < k {
					continue
				}
				tmp := longestSubstring(subStr, k)
				if tmp > res {
					res = tmp
				}
			}
			return res
		}
	}
	return len(s)
}

func main() {
	s := "abdbc"
	k := 2
	fmt.Println(longestSubstring(s, k))
}
