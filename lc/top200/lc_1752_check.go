package main

/*
检查数组是否经排序和轮转得到
给你一个数组 nums 。nums 的源数组中，所有元素与 nums 相同，但按非递减顺序排列。
如果nums 能够由源数组轮转若干位置（包括 0 个位置）得到，则返回 true ；否则，返回 false 。
源数组中可能存在 重复项 。
注意：我们称数组 A 在轮转 x 个位置后得到长度相同的数组 B ，当它们满足 A[i] == B[(i+x) % A.length] ，其中 % 为取余运算。


输入：nums = [3,4,5,1,2]
输出：true
解释：[1,2,3,4,5] 为有序的源数组。
可以轮转 x = 3 个位置，使新数组从值为 3 的元素开始：[3,4,5,1,2] 。

方案一：
1、找到第一个nums[i] < nums[i-1]的位置i
2、将i及以后的数组放到数组头部
3、判断数组是否完全符合非递减关系

方案二、
1、数组最多有一个位置满足 nums[i] < nums[i-1]
2、若存在1的情况，需满足 nums[0] >= nums[-1]
*/
func check(nums []int) bool {
	n := len(nums)
	i := 1
	for ; i < n; i++ {
		if nums[i] < nums[i-1] {
			break
		}
	}

	nums = append(nums[i:], nums[:i]...)

	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func check1752(nums []int) bool {
	t, n := 0, len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			t += 1
		}
		if t > 1 {
			return false
		}
	}
	return t == 0 || (t == 1 && nums[0] >= nums[n-1])
}
func main() {

}
