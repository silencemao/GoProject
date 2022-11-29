package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
生成交替二进制字符串的最少操作数

给你一个仅由字符 '0' 和 '1' 组成的字符串 s 。一步操作中，你可以将任一 '0' 变成 '1' ，或者将 '1' 变成 '0' 。
交替字符串 定义为：如果字符串中不存在相邻两个字符相等的情况，那么该字符串就是交替字符串。例如，字符串 "010" 是交替字符串，而字符串 "0100" 不是。
返回使 s 变成 交替字符串 所需的 最少 操作数。

输入：s = "0100"
输出：1
解释：如果将最后一个字符变为 '1' ，s 就变成 "0101" ，即符合交替字符串定义。


"10010100" 贪心会报错
*/
func minOperations(s string) int {
	res := 0
	ss := []byte(s)
	for i := 1; i < len(ss); i++ {
		if ss[i] == ss[i-1] {
			res += 1
			if ss[i] == '0' {
				ss[i] = '1'
			} else {
				ss[i] = '0'
			}
		}
	}
	return res
}

/*
结果只有01010... 或者 101010...两种情况
所以直接和这两种情况对比 不同的元素个数，取其中较小的即为最少操作数
*/
func minOperations1(s string) int {
	x := '0'
	cnt0, cnt1 := 0, 0
	for _, c := range s {
		if c != x {
			cnt0 += 1
		} else {
			cnt1 += 1
		}
		x ^= 1
	}
	return util.MinInt(cnt1, cnt0)
}

func main() {
	s := "10010100"
	fmt.Println(minOperations(s))
	fmt.Println(minOperations1(s))
}
