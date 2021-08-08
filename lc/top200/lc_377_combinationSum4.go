package main

import "fmt"

/*
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。

完全背包问题
*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

func help377(nums []int, target, cur int, res *int) {
	if cur > target {
		return
	}
	if cur == target {
		*res += 1
		return
	}
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		cur += nums[i]
		help377(nums, target, cur, res)
		cur -= nums[i]
	}
}

/*
超时
*/
func combinationSum4_2(nums []int, target int) int {
	res := 0
	help377(nums, target, 0, &res)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
	fmt.Println(combinationSum4_2(nums, target))
}
