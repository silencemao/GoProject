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

/*
栈，栈底始终存储最后一个没有被匹配的')'，栈中其它位置存储被遍历过的'('的索引
先用-1代替没有被匹配的')'，放入栈底
遇到'('，直接将索引加入栈中
遇到')' 首相将栈顶元素pop出去
若栈为空，则将当前')'的索引加入栈中
若栈不为空，则计算当前元素与栈顶元素的距离
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
			if len(stack) == 0 { // 栈中没有 与 当前')'匹配的元素 因为栈底是-1
				stack = append(stack, i)
			} else { // 被pop出去的那个元素是与当前')'匹配的元素
				res = util.MaxInt(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}
func main() {
	s := "((()"
	fmt.Println(longestValidParentheses(s))
}
