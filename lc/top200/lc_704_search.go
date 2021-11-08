package main

import (
	"fmt"
)

func search704(nums []int, target int) int {
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

	return res
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 2
	fmt.Println(search704(nums, target))
}
