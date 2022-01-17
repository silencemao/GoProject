package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。


*/
func maxProfitIII(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 5) // 不操作 第一次买 第一次卖 第二次买 第二次卖
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = util.MaxInt(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = util.MaxInt(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = util.MaxInt(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = util.MaxInt(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[len(prices)-1][4]
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfitIII(prices))
}
