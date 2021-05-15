package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)

func maxSubArray(nums []int) int {
	res, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = util.MaxInt(cur+nums[i], nums[i])
		res = util.MaxInt(res, cur)
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
