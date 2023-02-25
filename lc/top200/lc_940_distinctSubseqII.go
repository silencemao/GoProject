package main

import "fmt"

/*
给定一个字符串 s，计算 s 的 不同非空子序列 的个数。因为结果可能很大，所以返回答案需要对 10^9 + 7 取余 。
字符串的 子序列 是经由原字符串删除一些（也可能不删除）字符但不改变剩余字符相对位置的一个新字符串。
例如，"ace" 是 "abcde" 的一个子序列，但 "aec" 不是。

输入：s = "abc"
输出：7
解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
*/

func distinctSubseqII(s string) int {
	var mod int = 1e9 + 7
	res := 1
	preCnt := make([]int, 26)
	for _, c := range s {
		curAns := res
		res = ((res+curAns)%mod - preCnt[c-'a']%mod + mod) % mod
		preCnt[c-'a'] = res
	}

	return res - 1
}
func main() {
	s := "abd"
	fmt.Println(distinctSubseqII(s))
}
