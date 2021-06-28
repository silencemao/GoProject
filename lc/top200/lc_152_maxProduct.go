package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
连续子数组的最大乘积
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

1、因为可能会存在负数，而且负负得正，所以在计算的过程中也要考虑最小的乘积(最小的乘积*负数=正数)
2、因为是连续子数组，所以在更新minNum、maxNum的时候也要考虑当前数字nums[i]，这样才能保证子数组的连续性
*/
func maxProduct(nums []int) int {
	minNum, maxNum := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp1, tmp2 := minNum*nums[i], maxNum*nums[i]
		minNum = util.MinInt(util.MinInt(tmp1, tmp2), nums[i])
		maxNum = util.MaxInt(util.MaxInt(tmp1, tmp2), nums[i])
		res = util.MaxInt(util.MaxInt(minNum, maxNum), res)
	}

	return res
}
func main() {
	nums := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums))
}
