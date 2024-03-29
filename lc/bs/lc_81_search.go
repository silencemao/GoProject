package main

import "fmt"

/*
已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。
如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。
和第33题类似
*/

func search81(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		for l+1 < r && nums[l+1] == nums[l] { // 降低相等数据引起的影响
			l += 1
		}
		for r-1 > l && nums[r-1] == nums[r] {
			r -= 1
		}

		mid := l + (r-l)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] <= nums[r] {
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	target := 2
	fmt.Println(search81(nums, target))
}
