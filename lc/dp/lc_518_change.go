package main

import "fmt"

/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 凑成总金额为0的货币组合数为1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c] // 考虑货币c的组合数 dp[i-c] + 不考虑货币c的组合数dp[i]
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println(change(amount, coins))
}
