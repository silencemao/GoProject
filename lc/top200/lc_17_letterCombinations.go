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

func help17(res *[]string, digits, tmp string, ind int, letters map[int]string) {
	if ind >= len(digits) {
		*res = append(*res, tmp)
		return // 必须有这个return 否则 ind > len(digits)时，再从digits中取数会越界
	}

	digit := letters[int(digits[ind]-'0')]
	for i := 0; i < len(digit); i++ {
		tmp += string(digit[i])
		help17(res, digits, tmp, ind+1, letters)
		tmp = tmp[:len(tmp)-1]
	}
}

func letterCombinations2(digits string) []string {
	var res []string
	letters := make(map[int]string)
	letters[2] = "abc"
	letters[3] = "def"
	letters[4] = "ghi"
	letters[5] = "jkl"
	letters[6] = "mno"
	letters[7] = "pqrs"
	letters[8] = "tuv"
	letters[9] = "wxyz"
	help17(&res, digits, "", 0, letters)
	return res
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	fmt.Println(letterCombinations1(digits))
	fmt.Println(letterCombinations2(digits))
}
