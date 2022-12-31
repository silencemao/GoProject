package main

import "fmt"

/*
删除字符串两端相同字符后的最短长度
两端被删除的字符串长度不一定相等
*/
func minimumLength(s string) int {
	if len(s) < 2 || s[0] != s[len(s)-1] {
		return len(s)
	}
	i := 1
	for ; i < len(s); i++ {
		if s[i] != s[0] {
			break
		}
	}
	j := len(s) - 1
	for ; j >= i; j-- {
		if s[j] != s[len(s)-1] {
			break
		}
	}
	s = s[i : j+1]
	fmt.Println(s)
	return minimumLength(s)
}
func main() {
	s := "aabccabba"
	fmt.Println(minimumLength(s))
}
