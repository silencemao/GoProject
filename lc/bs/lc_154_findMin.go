package main

import "fmt"

/*
旋转数据寻找最小值，有重复元素
同第81 同第153题
*/
func findMin154(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		for l+1 < r && nums[l] == nums[l+1] {
			l += 1
		}
		for r-1 > l && nums[r] == nums[r-1] {
			r -= 1
		}
		mid := l + (r-l)>>1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = r - 1
		}
	}
	return nums[l]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	fmt.Println(findMin154(nums))
}
