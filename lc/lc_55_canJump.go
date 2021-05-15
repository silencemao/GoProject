package main

import (
	"fmt"
	"math"
)
/**
能否到达数组的最后一个位置
给定一个数组，全部为非负数，每一个位置的数代表在当前位置能向后跳跃的最大值
判断是否能跳跃到数组的最后一个位置

动态规划思路：
判断能否到达最后一个位置，也即判断跳力是否会为0
生成一个dp数组存储当前位置所拥有的跳力
每一个位置所拥有的跳力值为 max(上一个位置剩余的跳力值dp[i-1], 上一个位置新的跳力nums[i-1]) - 1
所以只要在每一个位置的跳力值不为0，就可以一直向后跳
思考：为什么当前位置的跳力值不是dp[i-1] + nums[i-1]
*/
func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i-1])) - 1)
		if dp[i] < 0 {
			return false
		}
	}
	return true
}
/*
贪心的思路，就是维护一个值reach，代表所能到达的最大值
如果坐标大于reach或者reach已经大于等于数组长度，则退出循环

reach的更新，max(当前位置的索引+当前位置能走的最远距离, reach)
*/
func canJump1(nums []int) bool {
	reach, n :=  0 , len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i > reach || reach >= n {
			break
		}
		reach = int(math.Max(float64(reach), float64(i+nums[i])))
	}
	return reach >= n
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
	fmt.Println(canJump1(nums))
}
