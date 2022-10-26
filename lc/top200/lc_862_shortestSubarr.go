package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
和至少为k的最短子数组

给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回 -1 。

子数组 是数组中 连续 的一部分。

前缀树+双端单调队列
*/
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	var q []int
	res := n + 1
	for i := 0; i <= n; i++ {
		for len(q) > 0 && preSum[i]-preSum[q[0]] >= k {
			res = util.MinInt(res, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && preSum[q[len(q)-1]] >= preSum[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if res == n+1 {
		res = -1
	}

	return res
}

func main() {
	nums := []int{2, -1, 2}
	k := 3
	fmt.Println(shortestSubarray(nums, k))
}
