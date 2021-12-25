package main

import "fmt"

/*
给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）

可以买卖多次，求最大收益
只要今天的价格 > 昨天的价格，就可以昨天买，今天卖，收益增加
*/
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}

	}
	return res
}
func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}
