package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)
/*
连续子数组的最大乘积，和第53题连续子数组的最大和一样
最大乘积 包括负数*负数 正数*正数
*/
func maxProduct(nums []int) int {
	curMin, curMax, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax := curMax * nums[i]
		tmpMin := curMin * nums[i]
		curMin = util.MinInt(util.MinInt(tmpMax, tmpMin), nums[i])
		curMax = util.MaxInt(util.MaxInt(tmpMax, tmpMin), nums[i])
		res = util.MaxInt(curMax, res)
	}
	return res
}

func main() {
	nums := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums))
}
