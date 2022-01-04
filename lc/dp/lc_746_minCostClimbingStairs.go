package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

输入：cost = [1,100,1,1,1,100,1,1,100,1]
输出：6
解释：你将从下标为 0 的台阶开始。
- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。

到达楼梯顶部是指跨越最后一个台阶(len(cost)-1)，并不到到达最后一个台阶
顶部的位置是len(cost)，所以dp的大小是len(cost)+1
*/
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	for i := 2; i < len(dp); i++ {
		dp[i] = util.MinInt(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	fmt.Println(cost)
	fmt.Println(dp)
	return dp[len(dp)-1]
}
func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
