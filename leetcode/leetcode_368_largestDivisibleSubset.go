package main

import (
	"fmt"
	"sort"
)

/*
给定数组nums,返回最长子数组subset，数组中的任意两个元素nums[i]%nums[j]=0 或 nums[j]%nums[i] = 0
*/
func largestDivisibleSubset(nums []int) []int {
	var res []int
	sort.Ints(nums)

	dp := make([]int, len(nums))
	parentId := make([]int, len(nums)) // 记录最长subset的各个元素的索引
	mx, mxId := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				parentId[i] = j // 记录索引，表示nums[j]%nums[i]=0
				if mx < dp[i] {
					mx = dp[i]
					mxId = i
				}
			}
		}
	}
	for i := 0; i < mx; i++ {
		res = append(res, nums[mxId])
		mxId = parentId[mxId]
	}

	return res
}

func findMaxLength(nums []int) {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				dp[i] = dp[j] + 1
				break
			}
		}
	}
	fmt.Println(dp)
}

func main() {
	nums := []int{1, 2, 4, 8, 7, 9, 10, 16}
	fmt.Println(largestDivisibleSubset(nums))
	findMaxLength(nums)
}
