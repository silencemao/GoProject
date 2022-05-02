package main

import "fmt"

/*
有序数组中的单一元素

给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
请你找出并返回只出现一次的那个数。
你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
*/
func singleNonDuplicate(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

func singleNonDuplicate1(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if mid%2 == 1 {
			if nums[mid] == nums[mid-1] { // 左侧没问题
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			if nums[mid] == nums[mid+1] { // 左侧没问题
				l = mid + 1
				//l = mid + 2  // 此处mid+1或mid+2都可以
			} else {
				r = mid
			}
		}
	}
	return nums[l]
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate1(nums))
	fmt.Println(singleNonDuplicate(nums))
}
