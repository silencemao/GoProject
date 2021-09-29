package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
打劫，不能同时打劫相邻的两家，否则会触发报警，如何在不触发报警的前提下打劫获得最大的收益

方案一：
直接利用一个与输入一样大的数组存储以i结尾可以获得的最大收益
dp[i]表示以i结尾可以获得多少，dp[i]由dp[i-2]+nums[i] 以及 dp[i-1]决定，所以转移方程为dp[i] = max(dp[i-2], dp[i-1])

方案二：
由于dp的长度太长，而且dp[i]仅有 dp[i-2]和dp[i-1]决定，因此可以使用两个变量代替 而不需要一个dp数组了
*/
func rob(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			dp[i] = util.MaxInt(nums[i], dp[i-1])
		} else {
			dp[i] = util.MaxInt(dp[i-2]+nums[i], dp[i-1])
		}
		res = util.MaxInt(dp[i], res)
	}
	return res
}

func rob1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	res, first, second := 0, nums[0], 0
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			second = util.MaxInt(first, nums[i])
		} else {
			first, second = second, util.MaxInt(first+nums[i], second)
		}
		res = util.MaxInt(second, res)
	}
	return res
}

func rob2(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	if size == 2 {
		return util.MaxInt(nums[0], nums[1])
	}
	first, second := nums[0], util.MaxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, util.MaxInt(first+nums[i], second)
	}
	return second
}

func main() {
	nums := []int{2, 7, 9, 3, 1, 7, 6, 8, 3, 10, 3}
	fmt.Println(rob(nums))
	fmt.Println(rob1(nums))
}
