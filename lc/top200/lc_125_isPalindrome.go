package main

import (
	"fmt"
	"strings"
)

// 验证字符串是否为回文串，只考虑字母和数字
func isPalindrome(s string) bool {
	s = strings.ToLower(s) // 全部变为小写
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isalnum(s[l]) { // 此处用for最好，如果是if的话，比较是否相等的时候也会有问题
			l++
		}
		for l < r && !isalnum(s[r]) {
			r--
		}
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
