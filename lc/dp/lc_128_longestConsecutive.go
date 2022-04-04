package main

import (
	"fmt"
	"sort"
)

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
// dp O(nlog(n))
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	dp := make([]int, len(nums)+1)
	res := 0
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
		} else if nums[i] == nums[i-1]+1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func longestConsecutive1(nums []int) int {
	var res int
	set := make(map[int]bool, 0)
	for _, num := range nums {
		set[num] = true
	}

	for _, num := range nums {
		if set[num-1] { // num-1存在set中，则以num-1为起点的长度最长
			continue
		}
		curNum := num
		curLength := 1

		for set[curNum+1] {
			curNum += 1
			curLength += 1
		}
		if curLength > res {
			res = curLength
		}
	}
	return res
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive1(nums))
}
