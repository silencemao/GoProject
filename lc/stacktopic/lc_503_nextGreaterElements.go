package main

import "fmt"

/*

给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。
数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

循环数组 采用2*n-1 单调栈
*/
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	stack := make([]int, 0)
	n := len(nums)
	for i := 0; i < 2*n-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = nums[i%n]
		}
		stack = append(stack, i%n)
	}
	return res
}

func nextGreaterElements1(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	n := len(nums)
	for i := 2*n - 1; i >= 0; i-- { // 相当于把数组扩充了一下，实际上没有扩充
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i%n] { // 注意等号，
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i%n] = nums[stack[len(stack)-1]]
		} else {
			res[i%n] = -1
		}
		stack = append(stack, i%n)
	}
	return res
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
	fmt.Println(nextGreaterElements1(nums))
}
