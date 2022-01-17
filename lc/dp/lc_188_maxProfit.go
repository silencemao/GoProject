package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfitIV(k int, prices []int) int {
	dp := make([][][]int, len(prices))
	k = util.MinInt(len(prices)/2, k)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2) // [持有, 不持有]
			dp[i][j][0] = -prices[0]
		}
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = util.MaxInt(dp[i-1][j-1][1]-prices[i], dp[i-1][j][0])
			dp[i][j][1] = util.MaxInt(dp[i-1][j][0]+prices[i], dp[i-1][j][1])
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1][k][1]
}

/*
使用二维数组 dp[i][j] ：第i天的状态为j，所剩下的最大现金是dp[i][j]
*/
func maxProfitIV2(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1) // 不操作 第一次买入 第一次卖出 第二次买入 第二次卖出...
		for j := 1; j < 2*k+1; j += 2 {
			dp[i][j] = -prices[0]
		}
	}
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = util.MaxInt(dp[i-1][j]-prices[i], dp[i-1][j+1])
			dp[i][j+2] = util.MaxInt(dp[i-1][j+1]+prices[i], dp[i-1][j+2])
		}
	}
	return dp[len(prices)-1][2*k]
}

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	fmt.Println(maxProfitIV(k, prices))
}
