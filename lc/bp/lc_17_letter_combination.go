package main

import "fmt"


func help17(res *[]string, tmp, digits string, ind int, set map[int]string) {
	if len(tmp) == len(digits) {
		*res = append(*res, tmp)
		return
	}
	digit := int(digits[ind] - '0')
	str := set[digit]
	for i := 0; i < len(str); i++ {
		tmp += string(str[i])
		help17(res, tmp, digits, ind+1, set)
		tmp = tmp[:len(tmp)-1]
	}
}

/*
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

*/
func letterCombinations(digits string) []string {
	set := make(map[int]string)
	set[0] = ""
	set[1] = ""
	set[2] = "abc"
	set[3] = "def"
	set[4] = "ghi"
	set[5] = "jkl"
	set[6] = "mno"
	set[7] = "pqrs"
	set[8] = "tuv"
	set[9] = "wxyz"

	var res []string
	help17(&res, "", digits, 0, set)
	return res

}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))

}
