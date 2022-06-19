package main

import "fmt"

/*
寻找重复数  https://leetcode-cn.com/problems/find-the-duplicate-number
给定一个包含n + 1 个整数的数组nums ，其数字都在 1 到 n之间（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。
你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。

和217、26、136一起看

快慢指针，
[1, 3, 4, 2]
0->1
1->3
2->4
3->2  生成链表 0-1-3-2-4-nil  无环

[1, 3, 4, 2, 2]
0->1
1->3
2->4
3->2
4->2 生成链表 0-1-3-2-4-2  有环
*/
func findDuplicate(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	l, f := 0, 0
	l, f = nums[l], nums[nums[f]] // f快速进入环中，在环内不断的走
	for l != f {
		l = nums[l]
		f = nums[nums[f]]
	}
	pre1, pre2 := 0, l
	for pre1 != pre2 {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}
	return pre1
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
