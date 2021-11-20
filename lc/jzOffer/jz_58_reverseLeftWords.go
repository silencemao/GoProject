package main

import "fmt"
/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，
该函数将返回左旋转两位得到的结果"cdefgab"。
*/
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	if n >= len(s) {
		n = n % len(s)
	}
	for i, j := 0, n-1; i < j; i, j = i + 1, j - 1 {
		b[i], b[j] = b[j], b[i]
	}
	for i, j := n, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	s = string(b)
	return s
}

func main() {
	s := "abcdefg"
	n := 7
	fmt.Println(reverseLeftWords(s, n))
}
