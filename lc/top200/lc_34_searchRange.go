package main

import "fmt"

func searchRange(nums []int, target int) []int {
	var res []int
	leftIndex := searchHelp(nums, true, target)
	rightIndex := searchHelp(nums, false, target) - 1
	if 0 <= leftIndex && leftIndex <= rightIndex && rightIndex < len(nums) && nums[leftIndex] == nums[rightIndex] && nums[leftIndex] == target {
		res = []int{leftIndex, rightIndex}
	} else {
		res = []int{-1, -1}
	}

	return res
}

func searchHelp(nums []int, left bool, target int) int {
	l, r := 0, len(nums)-1
	ans := len(nums) // 兼容nums只有一个数的时候
	for l <= r {
		mid := l + (r-l)>>1
		if (left && nums[mid] >= target) || nums[mid] > target {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

func main() {
	nums := []int{1}
	target := 1
	fmt.Println(searchRange(nums, target))
}
