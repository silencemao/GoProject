package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
买卖股票，只能进行一次交易
*/
// 超时
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			if prices[i] > prices[j] {
				res = util.MaxInt(prices[i]-prices[j], res)
			}
		}
	}
	return res
}

/*
dp[i]表示以i结尾的最大收益
*/
func maxProfit1(prices []int) int {
	dp := make([]int, len(prices))
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[i-1] { // 当前股票大于昨天的股票值，直接在昨天的收益上累加
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else { // 当前不大于昨天的股票价值，根据昨天和今天的差值计算今天的收益
			dp[i] = util.MaxInt(dp[i-1]-(-prices[i]+prices[i-1]), 0)
		}
		res = util.MaxInt(dp[i], res)
	}
	return res
}

/*
改为一维
*/
func maxProfit2(prices []int) int {
	res := 0
	curMax := 0
	for i := 1; i < len(prices); i++ {
		curMax = util.MaxInt(curMax+prices[i]-prices[i-1], 0)
		res = util.MaxInt(curMax, res)
	}
	return res
}

func main() {
	nums := []int{7, 6, 4}
	fmt.Println(maxProfit(nums))
	fmt.Println(maxProfit1(nums))
	fmt.Println(maxProfit2(nums))
}
