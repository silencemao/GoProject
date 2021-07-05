package main

import "fmt"

func letterCombinations(digits string) []string {
	//if len(digits) == 0 {
	//	return []string{}
	//}
	res := []string{""}
	dig2Char := make(map[int]string)
	dig2Char[2] = "abc"
	dig2Char[3] = "def"
	dig2Char[4] = "ghi"
	dig2Char[5] = "jkl"
	dig2Char[6] = "mno"
	dig2Char[7] = "pqrs"
	dig2Char[8] = "tuv"
	dig2Char[9] = "wxyz"

	for _, b := range []byte(digits) {
		var tmp []string
		chars := dig2Char[int(b-'0')]
		for i := 0; i < len(res); i++ {
			for j := 0; j < len(chars); j++ {
				tmp = append(tmp, res[i]+string(chars[j]))
			}
		}
		res = append([]string{}, tmp...)
	}

	return res
}

func dfs17(digits string, dig2Char map[int]string, res *[]string, tmp string, ind int) {
	if len(tmp) == len(digits) {
		*res = append(*res, tmp)
		return
	}
	chars := dig2Char[int(digits[ind]-'0')]
	for i := 0; i < len(chars); i++ { // 此处是对当前num对应的字符串进行搜索
		tmp += string(chars[i])
		dfs17(digits, dig2Char, res, tmp, ind+1) // 继续下一个num
		tmp = tmp[:len(tmp)-1]
	}
}

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dig2Char := make(map[int]string)
	dig2Char[2] = "abc"
	dig2Char[3] = "def"
	dig2Char[4] = "ghi"
	dig2Char[5] = "jkl"
	dig2Char[6] = "mno"
	dig2Char[7] = "pqrs"
	dig2Char[8] = "tuv"
	dig2Char[9] = "wxyz"

	res := []string{}
	dfs17(digits, dig2Char, &res, "", 0)
	return res
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	fmt.Println(letterCombinations1(digits))
}
