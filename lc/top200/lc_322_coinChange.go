package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回-1。

你可以认为每种硬币的数量是无限的。
https://leetcode-cn.com/problems/coin-change

像是一个背包问题，amount(背包) 使用硬币填满背包。
把背包的容量看成一个数组，分成amount份，记录填满背包每一份需要的硬币个数，硬币能否放进去背包中，需要比较coins[j]与i的大小
若能放进去，想使用最小的硬币个数的话，一种是状态是最好的，即将当前硬币放进去就能填满背包。所以我们可以在背包恰好剩余coins[i]的时候，将其放进去
此时背包内硬币的个数，为背包中已有i-coins[j]份硬币+1
即dp[i] = dp[i-coins[j]]+1
*/
func coinChange(coins []int, amount int) int {
	res := -1
	sort.Ints(coins)
	if amount == 0 {
		return 0
	}
	if coins[0] > amount {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = 1<<31 - 1
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = util.MinInt(dp[i], dp[i-coins[j]]+1)
			} else {
				break
			}
		}
	}
	if dp[amount] != 1<<31-1 {
		res = dp[amount]
	}
	return res
}
func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
