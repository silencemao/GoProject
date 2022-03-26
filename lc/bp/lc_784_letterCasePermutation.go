package main

import "fmt"

/*
给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。

返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
*/
func help784(res *[]string, s []byte, ind int) {
	*res = append(*res, string(s))
	for i := ind; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			s[i] -= 32
			help784(res, s, i+1)
			s[i] += 32
		}

		if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] += 32
			help784(res, s, i+1)
			s[i] -= 32
		}
	}
}

func letterCasePermutation(s string) []string {
	var res []string
	help784(&res, []byte(s), 0)
	return res
}

func main() {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))
}
