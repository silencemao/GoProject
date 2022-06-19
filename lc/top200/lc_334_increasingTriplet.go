package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 与第300题一起看
/*
dp会超时
*/
func increasingTriplet(nums []int) bool {
	dp := make([]int, len(nums))
	res := 0
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = util.MaxInt(dp[i], dp[j]+1)
			}
		}
		res = util.MaxInt(res, dp[i])
	}
	if res >= 3 {
		return true
	}
	return false
}

/*
给你一个整数数组nums ，判断这个数组中是否存在长度为 3 的递增子序列。
如果存在这样的三元组下标 (i, j, k)且满足 i < j < k ，使得nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
https://leetcode-cn.com/problems/increasing-triplet-subsequence

递增子序列，不要求序列的连续性。而且是三个，所以我们可以先找到长度为2的递增子序列[a, b]，然后再看后面的数有没有大于b的，
若存在大于b的数字则存在长度为3的递增子序列
若不存在可以更新 a b，让小的更小，才能找到满足条件的解
*/
func increasingTriplet1(nums []int) bool {
	first, second := 1<<31-1, 1<<31-1
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

/*
此题可以转换成更普适的情况，是否存在长度为n的递增子序列
*/
func increasingTriplet2(nums []int, n int) bool {
	if n > len(nums) || n < 1 {
		return false
	}
	res := make([]int, n-1)
	for i := range res {
		res[i] = 1<<31 - 1
	}
	for _, num := range nums {
		i := 0
		for ; i < n-1; i++ {
			if num < res[i] {
				res[i] = num
				break
			}
		}
		if i >= n-1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 3, 6, 1, 8}
	fmt.Println(increasingTriplet(nums))
	fmt.Println(increasingTriplet1(nums))
	fmt.Println(increasingTriplet2(nums, 7))
}
