package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

一：错误解法
nums中的数求和之后，直接根据res%2==0判断是否可以拆分成两个子数组，是错误的，比如[3,5]
二、背包问题
能否拆分成两个和相等的子集，相当于能否装满一个大小为res/2的背包(res为偶数)，每个数字只能用一次，所以是0-1背包
*/
func canPartition(nums []int) bool {
	res := 0
	for _, v := range nums {
		res += v
	}
	if res%2 == 1 {
		return false
	}
	size := res / 2
	dp := make([]int, size+1)
	for _, v := range nums {
		for j := size; j >= v; j-- {
			dp[j] = util.MaxInt(dp[j], dp[j-v]+v)
			if dp[j] == size {
				return true
			}
		}
	}
	return dp[size] == size
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
