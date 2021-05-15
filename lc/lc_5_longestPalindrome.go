package main

import "fmt"

func findPalindrome(s string, left, right int) string {
	bs := []byte(s)

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return string(bs[left+1 : right])
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		tmp := findPalindrome(s, i, i)
		if len(tmp) > len(res) {
			res = tmp
		}
		tmp = findPalindrome(s, i, i+1)
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func main() {
	s := "ac"

	fmt.Println(longestPalindrome(s))
}
