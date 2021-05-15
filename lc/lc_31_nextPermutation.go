package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	fmt.Println(nums)
	if len(nums) < 2 {
		return
	}
	var i int
	var flag bool = false
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			flag = true
			break
		}
	}
	if flag {
		var j int  // 从i位置向后找，找到i后面的数中最小的那个，但是是大于nums[i]的
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	i++
	for j := len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)
}


//func main() {
//	var nums = []int{5, 6 , 7, 5, 4, 3}
//	nextPermutation(nums)
//}