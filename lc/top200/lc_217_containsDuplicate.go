package main

import "fmt"

// 给定一个数组 判断是否存在重复的数字
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool, 0)
	set[nums[0]] = true
	for i := 1; i < len(nums); i++ {
		if set[nums[i]] {
			return true
		}
		set[nums[i]] = true
	}
	return false
}

/*
找出数组中重复的数字。
https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

因为数组的元素范围是0 ~ n-1，所以假如数组没有重复数字的话，每个位置i的值应该也是i，
1、即需要满足nums[i]=i
若不满足，则需要将nums[i]放到合适的位置，合适的位置在哪呢，即需要将其放在nums[nums[i]]处，
2、即需要nums[i] 与 nums[nums[i]]交换
交换完成之后，可以继续比较i位置的值是否满足条件1，若满足 i++ 若不满足 执行条件2
*/
func findRepeatNumber(nums []int) int {
	i := 0
	for true {
		if nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}
