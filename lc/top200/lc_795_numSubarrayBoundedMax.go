package main

import "fmt"

/*
给你一个整数数组 nums 和两个整数：left 及 right 。找出 nums 中连续、非空且其中最大元素在范围[left, right] 内的子数组，并返回满足条件的子数组的个数。
生成的测试用例保证结果符合 32-bit 整数范围。

输入：nums = [2,1,4,3], left = 2, right = 3
输出：3
解释：满足条件的三个子数组：[2], [2, 1], [3]
同第907题
*/
/*
单调递增的栈
1、以i为中心轴，向左找第一个nums[i]大的元素索引
2、向右找到第一个比nums[i]大的元素索引
3、i左侧数组长度 * i右侧数组长度，表示以i为最大值的子数组个数

*/

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	res := 0
	l, r := make([]int, len(nums)), make([]int, len(nums))
	for i := range l {
		l[i], r[i] = -1, len(nums)
	}
	var stack []int
	for i := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			r[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			l[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= left && nums[i] <= right {
			res += (r[i] - i) * (i - l[i])
		}
	}
	return res
}

/*
1、首先计算最大值在[0,right]区间内的子数组数量 a
2、再计算最大值在[0, left-1]区间内的子数组数量 b
3、返回a-b

子数组数量的计算res，遍历数组 ，用t记录子数组的长度，若num小于最大阈值 则t+=1，若大于最大阈值则t=0，res += t
*/
func numSubarrayBoundedMax1(nums []int, left int, right int) int {
	res := help795(right, nums) - help795(left-1, nums)
	return res
}
func help795(thresh int, nums []int) int {
	res, cnt := 0, 0
	for _, num := range nums {
		if num <= thresh {
			cnt += 1
		} else {
			cnt = 0
		}
		res += cnt
	}
	return res
}

func main() {
	nums := []int{2, 9, 2, 5, 6}
	left, right := 2, 8
	fmt.Println(numSubarrayBoundedMax1(nums, left, right))
	fmt.Println(nums)
	fmt.Println(numSubarrayBoundedMax(nums, left, right))
}
