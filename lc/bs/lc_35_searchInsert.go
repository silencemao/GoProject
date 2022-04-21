package main

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
*/
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := 0
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
			res = mid
		} else {
			l = mid + 1
			res = l
		}
	}
	return res
}

func searchInsert1(nums []int, target int) int {
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
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	res = l
	return res
}

func main() {

}
