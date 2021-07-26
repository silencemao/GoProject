package main

import "fmt"

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

*/
// 这种方法 对于 s='aa' t='aa' 不适用
func minWindow(s string, t string) string {
	res := ""

	need := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	var start, valid int

	for _, c := range t {
		need[byte(c)]++
	}
	size := 1<<31 - 1
	l, r := 0, 0
	for r < len(s) {
		c := s[r]
		r++
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for valid == len(t) {
			if r-l < size {
				start = l
				size = r - l
			}
			d := s[l]
			l++
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if size != 1<<31-1 {
		res = s[start : start+size]
	}

	return res
}

func minWindow1(s string, t string) string {
	res := ""

	need := make(map[byte]int, 128)
	for _, c := range t {
		need[byte(c)]++
	}
	start, valid, l := -1, 0, 0
	size := 1<<31 - 1

	for i := 0; i < len(s); i++ {
		need[byte(s[i])] -= 1
		if need[byte(s[i])] >= 0 { // s[i]为t中字符，所以valid记录一下 对于t中存在重复字符，也只会记录重复出现的次数，不会偏多也不会偏少
			valid++
		}
		for valid == len(t) { // 收缩窗口
			if i-l+1 < size {
				size = i - l + 1
				start = l
			}
			need[byte(s[l])]++        //
			if need[byte(s[l])] > 0 { // 此处是>0，对于t中出现的字符，need[s[l]] > 0 对于t中未出现的字符，need[s[l]]=0
				valid--
			}
			l++
		}
	}
	if start != -1 {
		res = s[start : start+size]
	}
	return res
}

func main() {
	s := "aa"
	t := "aa"
	fmt.Println(minWindow1(s, t))
}
