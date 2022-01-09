package main

import "fmt"

func dfs494(nums []int, sum, target, ind int, res *int) {

	if sum == target {
		*res += 1
	}

	for i := ind; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		sum += nums[i]
		dfs494(nums, sum, target, i+1, res)
		sum -= nums[i]
	}
}

func findTargetSumWays(nums []int, target int) int {
	res, sum := 0, 0
	for _, v := range nums {
		sum += v
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	target = (sum + target) >> 1
	dfs494(nums, 0, target, 0, &res)
	return res
}

func findTargetSumWays1(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	target = (sum + target) >> 1
	dp := make([]int, target+1) // dp[i]表示装满i的背包有多少种方法
	dp[0] = 1
	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] += dp[j-v]
		}
	}

	return dp[target]
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays1(nums, target))
}
