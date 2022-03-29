package main

import "fmt"

/*
给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10。

输入: 2
输出: 91
解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。
https://leetcode-cn.com/problems/count-numbers-with-unique-digits/solution/dong-tai-gui-hua-qing-xi-yi-dong-dai-ma-9lxva/
*/
func countNumbersWithUniqueDigits2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 9
	res := dp[1]
	for i := 2; i <= n; i++ { // 不包含0的时候
		dp[i] = dp[i-1] * (9 - (i - 1))
		res += dp[i]
	}
	for i := 2; i <= n; i++ { // 包含0的时候
		res += dp[i-1] * (i - 1)
	}
	return res + 1
}

func main() {
	n := 2
	//countNumbersWithUniqueDigits(n)
	fmt.Println(countNumbersWithUniqueDigits2(n))
}
