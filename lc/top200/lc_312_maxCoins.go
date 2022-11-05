package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
戳气球
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组nums中。
现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。
这里的 i - 1 和 i + 1 代表和i相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

求所能获得硬币的最大数量。

示例 1：
输入：nums = [3,1,5,8]
输出：167
解释：
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167

*/
func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	nums = append(append([]int{1}, nums...), 1)
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := nums[i] * nums[k] * nums[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = util.MaxInt(dp[i][j], sum)
			}
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return dp[0][n+1]
}

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}
