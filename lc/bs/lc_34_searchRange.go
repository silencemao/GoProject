package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
*/
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	l := help(nums, target, true)
	r := help(nums, target, false)
	if l >= 0 && r >= 0 && nums[l] == target && nums[r] == target {
		res[0] = l
		res[1] = r
	}
	return res
}

func help(nums []int, target int, isLeft bool) int {
	l, r := 0, len(nums)-1
	res := -1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target { // 此处不直接退出，而是根据找左/右做出不同的调整
			res = mid
			if isLeft {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return res
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}
