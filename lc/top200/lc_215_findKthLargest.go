package main

import (
	"fmt"
)

func partition(nums []int, left, right int) int {
	tmp := nums[left]
	for left < right {
		for left < right && nums[right] <= tmp {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= tmp {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = tmp
	return left
}

func helper(nums []int, left, right, k int) int {
	i := partition(nums, left, right)
	cnt := i - left + 1
	if cnt == k {
		return nums[i]
	} else if cnt > k {
		return helper(nums, left, i-1, k)
	} else {
		return helper(nums, i+1, right, k-cnt)
	}
}
func findKthLargest(nums []int, k int) int {
	res := helper(nums, 0, len(nums)-1, k)
	return res
}

func main() {
	nums := []int{9, 8, 0, 1, 4, 3, 9, 6, 5, 3}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}
