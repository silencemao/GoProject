package main

import "fmt"

/*
1、标准写法，找想等
2、第一个与目标值相等--左边界
3、最后一个与目标值相等--右边界
4、第一个大于等于目标值
5、最后一个小于等于目标值
*/

func binarySearch1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

//left
func binarySearch2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			r = mid - 1
		}
	}
	return -1
}

//right
func binarySearch3(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			l = mid + 1
		}
	}
	return -1
}

//first >=
func binarySearch4(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

//last <=
func binarySearch5(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] <= target {
			if mid+1 == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 1, 3, 3, 3, 4, 5, 5, 5, 5, 7, 7, 9}
	fmt.Println(nums)
	fmt.Println(binarySearch1(nums, 5))
	fmt.Println(binarySearch2(nums, 5))
	fmt.Println(binarySearch3(nums, 5))
	fmt.Println(binarySearch4(nums, 2))
	fmt.Println(binarySearch5(nums, 2))
}
