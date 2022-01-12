package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。

*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			dp[j] = util.MinInt(dp[j], dp[j-v]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{3}
	amount := 2
	fmt.Println(coinChange(coins, amount))
}