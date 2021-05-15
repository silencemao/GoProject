package main

import "fmt"

func sortHelp(nums []int, left, right int) {
	if left < right {
		tmp := nums[left]
		l, r := left, right
		for l < r {
			for l < r && nums[r] >= tmp {
				r--
			}
			if l < r {
				nums[l] = nums[r]
			}
			for l < r && nums[l] <= tmp {
				l++
			}
			if l < r {
				nums[r] = nums[l]
			}
		}
		nums[l] = tmp
		sortHelp(nums, left, l-1)
		sortHelp(nums, l+1, right)
	}
}
func sortColors(nums []int) {
	sortHelp(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func main() {
	nums := []int{2, 1, 0, 1, 2, 0, 9, 8, 7, 8, 10}
	sortColors(nums)
}
