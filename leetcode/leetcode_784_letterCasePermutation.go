package main

import (
	"unicode"
)
// 给定字符串"A1b2" 将字符串中的字母小写变大写，大写变小写，共有多少种组合
func upperLowerTransfer(s byte) byte {
	if s >= 'a' && s <= 'z' {
		return s - 32
	}
	return s + 32
}

func letterDFS(S string, index int, res *[]string) {
	if index == len(S) {
		*res = append(*res, S)
		return
	}
	var tmp string
	if unicode.IsLetter(rune(S[index])) {
		if unicode.IsLetter(rune(S[index])) {
			tmp = S[:index] + string(upperLowerTransfer(S[index])) + S[index+1:]
		} else {
			tmp = S
		}
		letterDFS(tmp, index+1, res)
		if unicode.IsLetter(rune(S[index])) {
			tmp = S[:index] + string(upperLowerTransfer(S[index])) + S[index+1:]
		} else {
			tmp = S
		}
	}
	letterDFS(S, index+1, res)
}

func letterCasePermutation(S string) []string {
	var res []string
	letterDFS(S, 0, &res)
	return res
}

func letterCasePermutation1(S string) []string {
	var res []string
	res = append(res, S)

	for i := 0; i < len(S); i++ {
		if unicode.IsLetter(rune(S[i])) {
			size := len(res)
			for j := 0; j < size; j++ {
				str := res[j]
				str = str[:i] + string(upperLowerTransfer(S[i])) + S[i+1:]
				res = append(res, str)
			}
		}
	}
	return res
}

//func main() {
//	S := "a1b2"
//	fmt.Println(letterCasePermutation1(S))
//}