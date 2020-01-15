package main

import (
	"fmt"
	"unicode"
)

/*
给定一个字符串，由字母和其他字符组成，将字母进行反转，其它字母的位置不动
abc-def
fed-cba
*/
func reverseOnlyLetters(S string) string {
	if len(S) > 1 {
		left, right := 0, len(S) - 1
		c := []byte(S)
		for left < right {
			if !unicode.IsLetter(rune(c[left])) {
				left++
			} else if !unicode.IsLetter(rune(c[right])){
				right--
			} else {
				c[left], c[right] = c[right], c[left]
				left++
				right--
			}
		}
		S = string(c)
	}
	return S
}

func main() {
	S := "abc-def"
	fmt.Print(reverseOnlyLetters(S))
}
