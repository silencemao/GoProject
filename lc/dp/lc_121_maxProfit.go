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

func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices)) // dp[i][0]表示当天持有股票，dp[i][1]表示当天不持有股票
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = util.MaxInt(dp[i-1][0], -prices[i]) // 此处不能写成dp[i][0] = util.MaxInt(dp[i-1][0], dp[i-1][1]-prices[i]) 因为只能买卖一次
		dp[i][1] = util.MaxInt(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(dp)-1][1]
}

//  空间优化
func maxProfit4(prices []int) int {
	dp := make([][]int, 2) // dp[i][0]表示当天持有股票，dp[i][1]表示当天不持有股票
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	size := len(prices)
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = util.MaxInt(dp[(i-1)%2][0], -prices[i])
		dp[i%2][1] = util.MaxInt(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
		//dp[i][0] = util.MaxInt(dp[i-1][0], -prices[i]) // 此处不能写成dp[i][0] = util.MaxInt(dp[i-1][0], dp[i-1][1]-prices[i]) 因为只能买卖一次
		//dp[i][1] = util.MaxInt(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[(size-1)%2][1]
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
	fmt.Println(maxProfit1(nums))
	fmt.Println(maxProfit2(nums))
	fmt.Println(maxProfit3(nums))
	fmt.Println(maxProfit4(nums))
}
