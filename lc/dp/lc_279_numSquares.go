package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = n
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = util.MinInt(dp[j], dp[j-i*i]+1)
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	n := 13
	fmt.Println(numSquares(n))
}
