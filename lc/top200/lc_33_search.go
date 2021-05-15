package main

import "fmt"

func searchRotate(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if nums[l] == nums[r] && target != nums[l] {
		return -1
	}

	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] { //左半边有序 此处小于等于很关键 mid位置不可能等于target了，但是l位置有可能
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右半边有序
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1}
	fmt.Println(searchRotate(nums, 1))
}
