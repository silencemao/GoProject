package main

import "fmt"

/*
生成有效的括号
递归
*/
func generateHelp(res *[]string, tmp string, l, r int) {
	if l > r {
		return
	}
	if l == 0 && r == 0 {
		*res = append(*res, tmp)
	} else {
		if l > 0 {
			generateHelp(res, tmp+"(", l-1, r)
		}
		if r > 0 {
			generateHelp(res, tmp+")", l, r-1)
		}
	}
}
func generateParenthesis(n int) {
	var res []string
	tmp := ""
	l, r := n, n
	generateHelp(&res, tmp, l, r)
	fmt.Println(res)
}

func main() {
	generateParenthesis(3)
}
