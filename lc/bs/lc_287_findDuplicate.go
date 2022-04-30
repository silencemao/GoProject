package main

import "fmt"

/*
给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回这个重复的数 。


寻找重复数，类似于找链表的环
// 0 nums[0], nums[nums[0]] nums[nums[nums[0]]].... 真正的头结点是0
*/
func findDuplicate(nums []int) int {
	s, f := nums[0], nums[nums[0]]
	for s != f {
		s, f = nums[s], nums[nums[f]]
	}
	s = 0 // 此处注意不能是nums[0]，因为头结点是0
	for s != f {
		s, f = nums[s], nums[f]
	}
	return s
}

// 二分查找
func findDuplicate1(nums []int) int {
	l, r := 1, len(nums)-1 // 数组长度为n+1,数字范围是1-n，此时的l r 不是索引，而是数值
	for l < r {
		mid := l + (r-l)>>1
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt += 1
			}
		}
		if cnt > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate1(nums))
}
