package main

import (
	"fmt"
	"math"
)

/*
给定一个数组nums[i]，计算由(i, nums[i]) (j, nums[j])组成的最大的矩形面积
*/
func maxArea(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			tmp := (i - j) * int(math.Min(float64(nums[i]), float64(nums[j])))
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

func maxArea2(nums []int) int {
	res := 0
	left, right := 0, len(nums)-1
	for left < right {
		tmp := (right - left) * int(math.Min(float64(nums[left]), float64(nums[right])))
		if tmp > res {
			res = tmp
		}
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
	fmt.Println(maxArea2(nums))
}
