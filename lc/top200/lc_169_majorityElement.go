package main

import (
	"fmt"
	"sort"
)

// 出现次数超过一半的元素，排序之后 中间位置的数一定是出现次数最多的元素
func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func help169(nums []int, left, right int) int {
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
		return l
	}
	return -1
}

func majority(nums []int, left, right int) int {
	mid := help169(nums, left, right)
	if mid == len(nums)/2 {
		return nums[mid]
	} else if mid > len(nums)/2 {
		return majority(nums, left, mid-1)
	} else {
		return majority(nums, mid+1, right)
	}
}

// 快排法 超时
func majorityElement1(nums []int) int {
	return majority(nums, 0, len(nums)-1)
}

// 投票法
func majorityElement2(nums []int) int {
	cnt := 0
	res := 0
	for _, num := range nums {
		if cnt == 0 {
			res = num
		}
		if num == res {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement2(nums))
}
