package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个整数数组prices，其中第i个元素代表了第i天的股票价格 ；整数fee 代表了交易股票的手续费用。
你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
返回获得利润的最大值。
注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

带有手续费的多次交易，和多次交易122题类似
*/
func maxProfitVI(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2) // [持有 不持有]
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = util.MaxInt(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = util.MaxInt(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[len(prices)-1][1]
}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfitVI(prices, fee))
}
