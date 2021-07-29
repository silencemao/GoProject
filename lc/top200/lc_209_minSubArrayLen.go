package main

import "fmt"

/*
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0

从nums中找出连续的子串，使子串的和>=target，返回最短的子串长度
https://leetcode-cn.com/problems/minimum-size-subarray-sum

滑动窗口，从左边开始遍历nums，累加nums[i]，直到valid>=target，记录子串长度，然后收缩左边届l
*/
func minSubArrayLen(target int, nums []int) int {
	res := 1<<31 - 1

	l, valid := 0, 0 // 左边界，nums求和
	for r := 0; r < len(nums); r++ {
		valid += nums[r]

		for valid >= target {
			//if valid == target {
			if r-l+1 < res {
				res = r - l + 1
				//}
			}
			valid -= nums[l] // 收缩左边界
			l++
		}
	}
	if res == 1<<31-1 {
		res = 0
	}

	return res
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 11
	fmt.Println(minSubArrayLen(target, nums))
}
