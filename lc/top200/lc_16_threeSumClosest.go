package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
最接近的三数之和
*/
func threeSumClosest(nums []int, target int) int {
	res := 1<<31 - 1
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if util.Abs(tmp-target) < util.Abs(res-target) {
				res = tmp
			}
			if tmp == target {
				return tmp
			} else if tmp > target {
				r -= 1
			} else {
				l += 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
