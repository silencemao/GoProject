package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func help343(res *[][]int, tmp []int, curSum, ind, n int) {
	if curSum > n || ind > n {
		return
	}
	if curSum == n {
		*res = append(*res, append([]int{}, tmp...))
		return
	}
	for i := ind; i < n; i++ {
		tmp = append(tmp, i)
		curSum += i

		help343(res, tmp, curSum, i, n) // 可以重复使用
		//help343(res, tmp, curSum, i+1, n) // 不可以重复使用

		tmp = tmp[:len(tmp)-1]
		curSum -= i
	}
}

func combine343(n int) {
	var res [][]int
	help343(&res, []int{}, 0, 1, n)
	fmt.Println(res)
}

/*
给定一个正整数n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36。

方案1：生成可以组成该整数n的所有组合，然后计算
方案2：动态规划，dp[i]表示 将i拆分为多个整数之后的最大乘积，当i>1之后，dp[i]表示的是至少两个整数的乘积
组成i的整数有很多种，[j,i-j] -> [[m, j-m], i-j] -> [[m, j-m], [n, i-j-n]] ...
所以我们可以遍历[1,i-1]之间的整数j，j可以与i-j构成i，j也可以与构成i-j的整数构成i(套娃😂 )，
上述两种情况的dp[i]分别为j*(i-j) 和 j*dp[i-j] dp[i-j]表示组成(i-j)的整数的最大乘积，
即递推公式为dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = util.MaxInt(util.MaxInt(j*(i-j), dp[i-j]*j), dp[i])
		}
	}
	fmt.Println(dp)
	return dp[n]
}
func main() {
	n := 10
	//combine343(n)
	fmt.Println(integerBreak(n))
}
