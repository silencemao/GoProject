package main

import "fmt"

/*
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array

将num[i]放到索引为num[i]-1的位置上，若nums[i]和索引为nums[i]-1位置上的数重复，则跳过，继续下一个

最后判断每个位置的数是否为索引 i+1
*/
func findDisappearedNumbers(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
