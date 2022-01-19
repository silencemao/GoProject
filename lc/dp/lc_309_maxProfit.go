package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个整数数组，其中第i个元素代表了第i天的股票价格
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
带有冷冻期的买卖股票
*/
func maxProfitV2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 4) // [买入, 保持卖出股票状态(已过冷冻期), 今天卖股票，冷冻期]
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = util.MaxInt(dp[i-1][0], util.MaxInt(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = util.MaxInt(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	size := len(prices)
	return util.MaxInt(util.MaxInt(dp[size-1][1], dp[size-1][2]), dp[size-1][3])
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfitV2(prices))
}
