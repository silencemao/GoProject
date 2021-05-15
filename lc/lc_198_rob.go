package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)
/*
打劫能获得的最多钱数
给定一个数组，表示每户拥有的钱，你可以打劫任意一家，但是不能打劫挨着的连续两家，这样能获得的最多的钱数

nums 1 3 1 3 100
dp   1 3 3 6 103

以上面的数组为例，我们以i位置为本次打劫的终点，计算i位置能获得的最大钱数
i位置能获得的钱由两个位置可以决定，一是(i-2)位置的钱数，二是(i-1)位置的钱数
为什么还能由i-1位置处的钱数决定呢？因为当i-2位置的钱数 与 i位置的钱数之和都小于(i-1)位置的钱数时，我们在i位置能获得的最大钱数是(i-1)位置处
的值，所以dp[i] = max(dp[i-2] + nums[i], dp[i-1])

*/
func rob(nums []int) int {
	mLen := len(nums)
	if mLen == 0 {
		return 0
	}
	if mLen == 1 {
		return nums[0]
	}
	dp := make([]int, mLen)

	dp[0] = nums[0]
	dp[1] = util.MaxInt(nums[0], nums[1])
	for i := 2; i < mLen; i++ {
		dp[i] = util.MaxInt(dp[i-2] + nums[i], dp[i-1])
	}
	return util.MaxInt(dp[mLen-1], dp[mLen-2])
}

func main() {
	nums := []int{1,3,1,3,100}
	fmt.Println(rob(nums))
}
