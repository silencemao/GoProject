package main

import (
	"fmt"
	"strings"
)

/*
判断字符串的两半是否相似

给你一个偶数长度的字符串 s 。将其拆分成长度相同的两半，前一半为 a ，后一半为 b 。

两个字符串 相似 的前提是它们都含有相同数目的元音（'a'，'e'，'i'，'o'，'u'，'A'，'E'，'I'，'O'，'U'）。注意，s 可能同时含有大写和小写字母。

如果 a 和 b 相似，返回 true ；否则，返回 false 。

输入：s = "book"
输出：true
解释：a = "bo" 且 b = "ok" 。a 中有 1 个元音，b 也有 1 个元音。所以，a 和 b 相似。
*/
func halvesAreAlike(s string) bool {
	size := len(s) / 2
	l, r := 0, size
	l1, l2 := 0, 0

	for l < size {
		if strings.ContainsAny(s[l:l+1], "aeiouAEIOU") {
			l1 += 1
		}
		if strings.ContainsAny(s[r:r+1], "aeiouAEIOU") {
			l2 += 1
		}
		l, r = l+1, r+1
	}
	fmt.Println(l1, l2)
	return l1 == l2
}

func main() {
	s := "book"
	fmt.Println(halvesAreAlike(s))
}
