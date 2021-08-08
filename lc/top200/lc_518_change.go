package main

import (
	"fmt"
)

/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

如果求组合数就是外层for循环遍历物品，内层for遍历背包。

如果求排列数就是外层for遍历背包，内层for循环遍历物品。

*/
func change(amount int, coins []int) int {
	res := 0
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= len(coins); i++ {
		val := coins[i-1]
		for j := val; j <= amount; j++ {
			dp[j] += dp[j-val]
		}
	}

	//for i := 1; i <= amount; i++ {
	//	for j := 0; j < len(coins); j++ {  // 会产生重复结果 3 [1, 1, 1] [1, 2] [2, 1]
	//		if coins[j] <= i {
	//			dp[i] += dp[i-coins[i]]
	//		} else {
	//			break
	//		}
	//	}
	//}

	res = dp[amount]
	return res
}

/*
错误 这是得到的排列个数，而本题求的是组合个数
按照下面的遍历方式 amount=3的时候，会先遍历到1，然后加上dp[3-1]，接着会再遍历到2 再加上dp[3-2] 产生了重复
*/
func change2(amount int, coins []int) int {
	res := 0
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			dp[i] += dp[i-coins[j]]
		}
	}
	res = dp[amount]
	return res
}

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
	fmt.Println(change2(amount, coins))
}

// 1 1 1 1 1、1 1 1 2、1 2 2、5
