package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数2、3 和/或5的正整数。
输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。

*/
func nthUglyNumber(n int) int {

	dp := make([]int, n+1)
	dp[1] = 1
	p1, p2, p3 := 1, 1, 1
	for i := 2; i <= n; i++ {
		num2, num3, num5 := dp[p1]*2, dp[p2]*3, dp[p3]*5
		tmp := util.MinInt(util.MinInt(num2, num3), num5)
		dp[i] = tmp
		if tmp == num2 {
			p1 += 1
		}
		if tmp == num3 {
			p2 += 1
		}
		if tmp == num5 {
			p3 += 1
		}
	}
	return dp[n]
}

func main() {
	n := 10
	fmt.Println(nthUglyNumber(n))
}
