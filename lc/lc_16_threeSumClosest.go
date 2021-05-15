package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个数组和target
返回数组内的三个数之和最接近target的值
最接近的三数之和
*/
func threeSumClosest(nums []int, target int) int {
	closet := nums[0] + nums[1] + nums[2]
	diff := int(math.Abs(float64(target - closet)))

	sort.Ints(nums)
	for i := 0; i < len(nums) - 1; i++ {
		left, right := i +1, len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			newDiff := int(math.Abs(float64(target - sum)))
			if newDiff < diff {
				diff = newDiff
				closet = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return closet
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}