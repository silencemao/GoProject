package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 可以进行k次交易
// dp[i][j][0] 在第i天第j次持有的收益 dp[i][j][1] 在第i天第j次不持有的收益
func maxProfitK(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	k = util.MinInt(k, len(prices)/2)
	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -prices[0] // 持有
		}
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = util.MaxInt(dp[i-1][j-1][1]-prices[i], dp[i-1][j][0]) // 持有 注意j-1 表示昨天的上一次状态
			dp[i][j][1] = util.MaxInt(dp[i-1][j][0]+prices[i], dp[i-1][j][1])
		}
	}

	return dp[len(prices)-1][k][1]
}

func f188(k int, prices []int) int {
	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -prices[i]
		}
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = util.MaxInt(dp[i-1][j-1][1]-prices[i], dp[i-1][j][0])
			dp[i][j][1] = util.MaxInt(dp[i-1][j][0]+prices[i], dp[i-1][j][1])
		}
	}
	return dp[len(prices)-1][k][1]
}

func main() {
	nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	k := 3
	fmt.Println(maxProfitK(k, nums))
	fmt.Println(f188(k, nums))
}
