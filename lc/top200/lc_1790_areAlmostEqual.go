package main

import "fmt"

/*
给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。

如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。
*/
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	diff, c1, c2, flag := 0, s1[0], s2[0], false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff += 1
			if diff > 2 {
				return false
			}
			if diff == 1 {
				c1, c2 = s1[i], s2[i]
			}
			if diff == 2 {
				flag = c1 == s2[i] && c2 == s1[i]
			}
		}

	}
	return flag
}

func main() {
	s1, s2 := "bank", "kanb"
	fmt.Println(areAlmostEqual(s1, s2))
}
