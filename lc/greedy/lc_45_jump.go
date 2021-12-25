package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = i
		for j := i - 1; j >= 0; j-- {
			if nums[j]+j >= i {
				dp[i] = util.MinInt(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}

func jump1(nums []int) int {
	res, curDistance, nextDistance := 0, 0, 0 // 结果，当前能到达的最远距离，下一步能到达的最远距离
	for i := 0; i < len(nums)-1; i++ {        // 最后一个位置不需要再跳了
		nextDistance = util.MaxInt(nextDistance, nums[i]+i) // 在[0, i]中能到达的最远距离
		if i == curDistance {                               // i走到了[0, i]之间某位置可以到达的最远距离，要进行下一跳了，并更新当前可以到达的最远距离
			curDistance = nextDistance
			res += 1
		}
	}
	return res
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
	fmt.Println(jump1(nums))
}
