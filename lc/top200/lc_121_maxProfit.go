package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 只能交易一次
func maxProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			res = util.MaxInt(res, prices[i]-prices[j])
		}
	}
	return res
}

func maxProfit1(prices []int) int {
	res := 0
	dp := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		dp[i] = util.MaxInt(dp[i-1]+prices[i]-prices[i-1], 0)
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
