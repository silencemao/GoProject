package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组[4,3]是该条件下的长度最小的子数组。
*/
// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	valid := 0
	l, res := 0, 1<<31-1
	for i := 0; i < len(nums); i++ {
		valid += nums[i]
		for valid >= target {
			res = util.MinInt(res, i-l+1)
			valid -= nums[l]
			l += 1
		}
	}
	return res
}

// 前缀树+二分法
func minSubArrayLen1(target int, nums []int) int {
	sums := make([]int, len(nums)+1)
	n := len(nums)
	res := n
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		s := target + sums[i-1]
		l, r := i, len(sums)-1
		if sums[r] < s { // 重要
			continue
		}
		for l < r {
			mid := l + (r-l)>>1
			if sums[mid] >= s {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if l != -1 {
			res = util.MinInt(res, l-(i-1))
		}
	}

	return res
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
	fmt.Println(minSubArrayLen1(target, nums))
}
