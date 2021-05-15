package main

import (
	"fmt"
	"strconv"
)

/*
1-26分别与a-z分别一一对应，给定一个数字表示的字符串，计算该字符串可以由几种字母组合表示
因为数字有一位数的也有两位数的，因此我们要考虑数字自身可以映射成一个字母以及与相邻数字组合组成字母两种情况
动态规划法：
以数字i为结尾时，字母组合个数由两种情况决定。1、数字i是否在1-9内，若在，则res[i] = res[i-1]
2、数字i-1与数字i的组合是否在10-26范围内，若在，则res[i] = res[i-2]
3、若上述两种情况都满足，则res[i]=res[i-1]+res[i-2]
*/
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	s = " " + s
	res := make([]int, len(s)+1)
	res[0] = 1
	for i := 1; i < len(s); i++ {
		a := s[i] - '0'
		b := (s[i-1]-'0')*10 + s[i] - '0'
		if 1 <= a && a <= 9 {
			res[i] = res[i-1]
		}
		if 10 <= b && b <= 26 {
			res[i] += res[i-2]
		}
	}
	return res[len(s)-1]
}

/*
深度优先搜索
*/
func isValid91(s string) bool {
	if len(s) == 0 || s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	if 1 <= num && num <= 26 {
		return true
	}
	return false
}
func dfs91(s string, res *int, ind int) {
	if ind >= len(s) {
		*res += 1
	}
	if ind+1 <= len(s) && isValid91(s[ind:ind+1]) {
		dfs91(s, res, ind+1)
	}
	if ind+2 <= len(s) && isValid91(s[ind:ind+2]) {
		dfs91(s, res, ind+2)
	}
}

func numDecodings1(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	res := 0
	dfs91(s, &res, 0)
	return res
}

func main() {
	s := "226"
	//fmt.Println(numDecodings(s))
	fmt.Println(numDecodings1(s))
}
