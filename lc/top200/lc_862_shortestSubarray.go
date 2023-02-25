package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func shortestSubArray(nums []int, k int) int {
	n, res := len(nums), len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	var q []int
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
	return res
}

func main() {
	nums := []int{2, -1, 2}
	k := 3
	fmt.Println(shortestSubArray(nums, k))
}
