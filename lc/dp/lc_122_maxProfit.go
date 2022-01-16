package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/
func maxProfitII(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ { // 贪心，只要今天价格高我就交易
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func maxProfitII2(prices []int) int {
	dp := make([][]int, 2) // dp[i][0]持有股票，不持有股票
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	size := len(prices)
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = util.MaxInt(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = util.MaxInt(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
	}
	return dp[(size-1)%2][1]
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfitII(prices))
	fmt.Println(maxProfitII2(prices))
}
