package main

import "fmt"

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

*/
func reverseStr(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(s); i += 2*k {
		l, r := i, i+k-1
		if r >= len(s) {
			r = len(s)-1
		}
		for l < r {
			ss[l], ss[r] = ss[r], ss[l]
			l++
			r--
		}
	}
	return string(ss)
}

func main() {
	s := "a"
	k := 2
	fmt.Println(reverseStr(s, k))
}
