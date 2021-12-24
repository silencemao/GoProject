package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 动态规划
func maxSubArray(nums []int) int {
	curSum, res := nums[0], nums[0]
	for _, v := range nums[1:] {
		curSum = util.MaxInt(curSum+v, v)
		res = util.MaxInt(res, curSum)
	}
	return res
}

// 贪心
func maxSubArray1(nums []int) int {
	res := -1 << 31
	curSum := 0
	for _, v := range nums {
		curSum += v
		if curSum > res {
			res = curSum
		}
		if curSum < 0 {
			curSum = 0
		}
	}
	return res
}

func main() {
	num := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(num))
}
