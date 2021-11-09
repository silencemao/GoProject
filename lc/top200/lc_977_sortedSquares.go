package main

import "fmt"
/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
*/
func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	k := r
	res := make([]int, len(nums))
	for l <= r {  // 一定要<= 否则最后k=0位置不会被赋值
		l2, r2 := nums[l]*nums[l], nums[r]*nums[r]
		if l2 < r2 {
			res[k] = r2
			r--
		} else {
			res[k] = l2
			l++
		}
		k--
	}
	return res
}
func main() {
	nums := []int{-4,-1,0,3,10}
	fmt.Println(sortedSquares(nums))
}
