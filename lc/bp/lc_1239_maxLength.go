package main

import "fmt"

/*
串联字符串的最大长度
给定一个字符串数组 arr，字符串 s 是将 arr的含有 不同字母 的子序列 字符串 连接 所得的字符串。
请返回所有可行解 s 中最长长度。
子序列 是一种可以从另一个数组派生而来的数组，通过删除某些元素或不删除元素而不改变其余元素的顺序。

输入：arr = ["cha","r","act","ers"]
输出：6
解释：可能的解答有 "chaers" 和 "acters"
给定字符串数组，返回可以组成的最大长度的字符串s，保证s中无重复字母
*/
func check1239(str string) bool {
	if len(str) > 26 {
		return false
	}
	charMap := make(map[byte]int, 0)
	for i := 0; i < len(str); i++ {
		charMap[str[i]] += 1
		if charMap[str[i]] > 1 {
			return false
		}
	}
	return true
}

func help1239(arr []string, curStr string, res *int, start int) {
	if len(curStr) > 26 {
		return
	}

	if len(curStr) > *res {
		*res = len(curStr)
	}
	for i := start; i < len(arr); i++ {
		if check1239(curStr + arr[i]) {
			curStr += arr[i]
			help1239(arr, curStr, res, i+1)
			curStr = curStr[:len(curStr)-len(arr[i])]
		}
	}
}

func maxLength(arr []string) int {
	res := 0
	help1239(arr, "", &res, 0)
	return res
}

func main() {
	arr := []string{"cha", "r", "act", "ers"}
	fmt.Println(maxLength(arr))
}
