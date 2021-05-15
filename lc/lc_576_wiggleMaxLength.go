package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)

func wiggleMaxLength(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}

	return util.MaxInt(up, down)
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(wiggleMaxLength(nums))
}
