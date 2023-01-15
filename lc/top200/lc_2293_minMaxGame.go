package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		n := len(nums) / 2

		for i := 0; i < n; i++ {
			if i%2 == 0 {
				nums[i] = util.MinInt(nums[2*i], nums[2*i+1])
			} else {
				nums[i] = util.MaxInt(nums[2*i], nums[2*i+1])
			}
		}
		nums = nums[:n]
	}
	return nums[0]
}

func main() {
	nums := []int{70, 38}
	fmt.Println(minMaxGame(nums))
}
