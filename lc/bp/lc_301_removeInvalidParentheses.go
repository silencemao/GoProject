package main

import "fmt"

/*
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

返回所有可能的结果。答案可以按 任意顺序 返回。

输入：s = "()())()"
输出：["(())()","()()()"]

*/
func help301(res *[]string, s string, l, r, ind int) {
	if l == 0 && r == 0 {
		if isValid301(s) {
			fmt.Println(s)
			*res = append(*res, s)
		}
		return
	}

	for i := ind; i < len(s); i++ {
		if i > ind && s[i] == s[i-1] {
			continue
		}
		//if l+r > len(s) {
		//	continue
		//}
		if l > 0 && s[i] == '(' {
			help301(res, s[:i]+s[i+1:], l-1, r, i) // 此处不需要i+1
		}
		if r > 0 && s[i] == ')' {
			help301(res, s[:i]+s[i+1:], l, r-1, i)
		}
	}

}

func isValid301(s string) bool {
	var stack []string

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, "(")
		}
		if s[i] == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func removeInvalidParentheses(s string) []string {
	var res []string
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			if l > 0 {
				l--
			} else {
				r++
			}
		}
	}
	help301(&res, s, l, r, 0)
	return res
}

func main() {
	s := ")("
	fmt.Println(removeInvalidParentheses(s))
}
