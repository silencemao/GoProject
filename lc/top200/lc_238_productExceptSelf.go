package main

import "fmt"

/*
给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。
https://leetcode-cn.com/problems/product-of-array-except-self

给定一个数组，
结果返回一个数组，数组i位置的值为其它位置的乘积，不包括i位置的值
两次遍历，第一次遍历i位置存储i左边数字的乘积，第二次遍历output[i] 与右侧的数字的乘积 相乘即可。
在第二次遍历的过程中，需要一个变量来记录右侧数字的垒乘，然后与i位置左边的乘积相乘
*/
func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	output[0] = 1
	for i := 1; i < len(nums); i++ {
		output[i] = output[i-1] * nums[i-1]
	}
	r := 1
	for i := len(nums) - 2; i >= 0; i-- {
		r *= nums[i+1]
		output[i] = r * output[i]
	}
	return output
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
