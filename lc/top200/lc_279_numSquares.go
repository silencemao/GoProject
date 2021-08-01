package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
完全平方数 (该数是一个整数，另一个整数的平方得来)

给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
https://leetcode-cn.com/problems/perfect-squares
此题就是找到最少由几个完全平方数可以组成给定的数n

可以转换为背包问题，将该数n看成一个大小为n的背包，我们可以计算背包的每一份最少可以由几个完全平方数组成
例如n=1 由一个完全平方数组成
n=2 由2个完全平方数组成
n=3 由3个完全平方数组成
组成n的完全平方数的范围肯定在 [1-n/2]的范围内（确切的说应该是[0, sqrt(n)]）
遍历j [1-n/2]，若该数的平方小于n，则j*j可以是组成n的一个完全平方数，而在n中要腾出j*j的空间，即n-j*j，
我们查看组成n-j*j最少需要几个完全平方数，即转移方程为dp[n]=min(dp[n-j*j], dp[n])

类似于322题 coin change
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = 1<<31 - 1
		for j := 1; j*j <= i; j++ {
			if j*j <= i {
				dp[i] = util.MinInt(dp[i-j*j]+1, dp[i])
			}
		}
	}
	return dp[n]
}

func main() {
	n := 12
	fmt.Println(numSquares(n))
}
