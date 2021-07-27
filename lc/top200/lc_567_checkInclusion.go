package main

import "fmt"

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的 子串 。

滑动窗口
*/
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int, 128)
	for _, b := range s1 {
		need[byte(b)] += 1
	}

	l, valid := 0, 0
	for i := 0; i < len(s2); i++ {
		need[s2[i]]--
		if need[s2[i]] >= 0 {
			valid++
		}

		for valid == len(s1) {
			if i-l+1 == len(s1) {
				return true
			}
			need[s2[l]]++
			if need[s2[l]] > 0 {
				valid--
			}
			l++
		}
	}

	return false
}
func main() {
	s1 := "ab"
	s2 := "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))
}
