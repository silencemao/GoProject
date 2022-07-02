package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
最长有效括号
给你一个只包含 '('和 ')'的字符串，找出最长有效（格式正确且连续）括号子串的长度。

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

*/
func longestValidParentheses(s string) int {
	res := 0
	var stack []int
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = util.MaxInt(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}
func main() {
	s := "()()"
	fmt.Println(longestValidParentheses(s))
}
