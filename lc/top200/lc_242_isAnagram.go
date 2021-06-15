package main

import "fmt"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
目的：判断两个字符串是否字母数量一致
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	set := make(map[int]int, 0)
	for i := 0; i < len(s); i++ {
		set[int(s[i]-'a')] += 1
		set[int(t[i]-'a')] -= 1
	}
	for i := range set {
		if set[i] != 0 {
			return false
		}
	}
	return true
}
func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
