package main

import "fmt"
/*
给定字符串s和t，如果t删除一部分字符之后和s相等，则返回true，否则返回false
*/
func isSubsequence(s, t string) bool {
	 res := false
	 if len(s) > len(t) {
	 	return res
	 }
	 indexS := 0
	 for i := 0; i < len(t); i++ {
	 	if t[i]==s[indexS] {
	 		indexS++
	 		if indexS==len(s) {
	 			return true
			}
		}
	 }
	 return res
}

func isSubsequence1(s, t string) bool {
	sLen, tLen := len(s), len(t)
	i, j := 0, 0

	for i < sLen && j < tLen {
		if s[i]==t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i==sLen
}
func main() {
	s := "adc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
	fmt.Println(isSubsequence1(s, t))
}
