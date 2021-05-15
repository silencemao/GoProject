package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)

func canJump(nums []int) bool {
	k := 0 // 表示能到达的的最远位置
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = util.MaxInt(i+nums[i], k)
		if k >= len(nums)-1 {
			return true
		}
	}
	return true
}

func main() {
	nums := []int{3, 0, 8, 2, 0, 0, 1}
	fmt.Println(canJump(nums))
}
