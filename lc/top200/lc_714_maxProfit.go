package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func maxProfitFee(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0          // 不持有
	dp[0][1] = -prices[0] // 持有
	for i := 1; i < len(prices); i++ {
		dp[i][0] = util.MaxInt(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = util.MaxInt(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func main() {
	nums := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfitFee(nums, fee))
}
