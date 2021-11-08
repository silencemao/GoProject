package main

import "fmt"
/*
给定无重复 排序数组nums 以及目标值target
若nums中存在target，返回位置
若不存在，返回插入位置
*/
func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	res := -1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid-1
		} else {
			l = mid+1
		}
	}

	res = l
	return res
}

func main() {
	nums := []int{1,3,5,6}
	target := 4
	fmt.Println(searchInsert( nums, target))
}
