package main

import "fmt"
/*
给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。

*/
func reverseWords(s string) string {
	b := []byte(s)
	slow, fast := 0, 0

	for len(b) > 0 && fast < len(b) && b[fast] == ' ' {  // 删除头部空格
		fast += 1
	}

	for ; fast < len(b); fast+=1  { // 删除单词间空格
		if fast > 0 && b[fast] == b[fast-1] && b[fast] == ' ' {
			continue
		}
		b[slow] = b[fast]
		slow += 1
	}

	if slow - 1 > 0 && b[slow-1] == ' ' { // 删除末尾空格
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}

	reverseStr151(&b, 0, len(b)-1)

	i, j := 0, 0
	for i < len(b) {
		j = i
		for ; j < len(b) && b[j] != ' '; {
			j += 1
		}
		reverseStr151(&b, i, j-1)
		i = j + 1
	}


	return string(b)
}

func reverseStr151(str *[]byte, i, j int) []byte {
	for i < j {
		(*str)[i], (*str)[j] = (*str)[j], (*str)[i]
		i++
		j--
	}
	return *str
}

func main() {
	s := "  hello  word "
	fmt.Println(reverseWords(s))
}
