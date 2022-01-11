package main

import "fmt"

/*
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
题目数据保证答案符合 32 位整数范围。
输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
求排列数
*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, v := range nums { // 先遍历物品 再遍历背包 只能得到组合数
		for j := v; j <= target; j++ {
			dp[j] += dp[j-v]
		}
	}
	return dp[target]
}

func combinationSum41(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ { // 先遍历背包 再遍历物品 得到排列数
		for _, v := range nums {
			if v <= i {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum41(nums, target))
}
