package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
第N个神奇数字
一个正整数如果能被 a 或 b 整除，那么它是神奇的。
给定三个整数 n ,a , b ，返回第 n 个神奇的数字。因为答案可能很大，所以返回答案对109+ 7 取模后的值。

输入：n = 1, a = 2, b = 3
输出：2
*/
func nthMagicalNumber(n int, a int, b int) int {
	const mod int = 1e9 + 7
	l, r := util.MinInt(a, b), n*util.MinInt(a, b)
	lcm := a / gcd(a, b) * b // 最小公倍数
	for l <= r {
		mid := l + (r-l)>>1
		cnt := mid/a + mid/b - mid/lcm //<=mid的数字中有多少个可以被 a 或者 b整除

		if cnt >= n { // cnt==n时
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return (r + 1) % mod
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	a, b, n := 2, 3, 4
	fmt.Println(gcd(a, b))
	fmt.Println(nthMagicalNumber(n, a, b))
}
