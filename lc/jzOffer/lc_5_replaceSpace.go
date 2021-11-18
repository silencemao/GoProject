package main

import "fmt"

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

输入：s = "We are happy."
输出："We%20are%20happy."

*/
func replaceSpace(s string) string {
	ssize, esize := 0, 0
	for i := range s {
		if s[i] == ' ' {
			esize += 1
		} else {
			ssize += 1
		}
	}
	tmp := make([]byte, ssize + esize*3)
	i, j := len(tmp)-1, len(s)-1
	for j >= 0 {
		if s[j] != ' ' {
			tmp[i] = s[j]
			i--
			j--
		} else {
			tmp[i] = '0'

			i--
			tmp[i] = '2'

			i--
			tmp[i] = '%'

			i--
			j--
		}
	}
	s = string(tmp)
	return s
}

func main() {
	s := "We are happy"
	fmt.Println(replaceSpace(s))
}
