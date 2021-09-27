package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 只能交易一次
func maxProfit(prices []int) int {
	res := 0
	// 以第i天结尾 交易一次 可以获得收益
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			res = util.MaxInt(res, prices[i]-prices[j])
		}
	}
	return res
}

func maxProfit1(prices []int) int {
	res := 0
	dp := make([]int, len(prices)) // dp[i]存储第i天卖出时时 可获得的最大收益
	for i := 1; i < len(prices); i++ {
		dp[i] = util.MaxInt(dp[i-1]+prices[i]-prices[i-1], 0) // 这种写法可以累加收益，即在昨天的基础上我是否还能获得更多收益
		res = util.MaxInt(dp[i], res)
	}
	return res
}

// 累加求和，
func maxProfit2(prices []int) int {
	res, curMax := 0, 0
	for i := 1; i < len(prices); i++ {
		curMax = util.MaxInt(curMax+prices[i]-prices[i-1], 0)
		res = util.MaxInt(curMax, res)
	}
	return res
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1(nums))
}
