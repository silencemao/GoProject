package main

import (
	"fmt"
	"GoProject/leetcode/util"
)

/*
最大连续子数组
给定一个数组，数组中数字有正有负，找出一段连续子数组，子数组中数字求和最大

本题中数组有正有负，并且要求子数组是连续的，为了保证连续性，所以有些负数也是要考虑进来的。
我们用curSum表示当前连续子数组的和，maxSum表示全局连续子数组的和。
首先 curSum, maxSum = nums[0], nums[0]
然后 遍历数组，当遇到正数，肯定是直接更新curSum = curSum + nums[i]
             当遇到负数，就要比较curSum + nums[i] 与 nums[i]的大小关系，二者取较大的数 curSum = max(curSum + nums[i], nums[i])
     为什么取较大的数，对于i之后的数字，curSum相当于起始值，起始值越大连续子数组的和才会越大，所以取较大的数
     所以不管遇到正数还是负数，都可以以curSum = max(curSum + nums[i], nums[i])来更新curSum

最后 更新maxSum
*/
func maxSubArray(nums []int) int {
	mLen := len(nums)
	curSum, maxSum := nums[0], nums[0]
	for i := 1; i < mLen; i++ {
		curSum = util.MaxInt(curSum + nums[i], nums[i])  // 向后看
		maxSum = util.MaxInt(curSum, maxSum)
	}
	return maxSum
}

func maxSubArray1(nums []int) int {
	mLen := len(nums)
	dp := make([]int, mLen)
	dp[0]=  nums[0]
	res := nums[0]
	for i := 1; i < mLen; i++ {
		dp[i] = nums[i] + util.MaxInt(dp[i-1], 0)   // 向前看，前面小于0，直接不要
		res = util.MaxInt(dp[i], res)
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray1(nums))
}
