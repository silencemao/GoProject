package main

import (
	"fmt"
	"sort"
)

/*
最长连续子数组
*/
func longestConsecutive(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	res = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			dp[i] = dp[i-1] + 1
		} else if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
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
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	dp := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			dp += 1
		} else if nums[i] == nums[i-1] {

		} else {
			dp = 1
		}
		if dp > res {
			res = dp
		}
	}
	return res
}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	set := make(map[int]int)
	sort.Ints(nums)
	set[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			continue
		}
		if cnt, ok := set[nums[i]-1]; ok {
			set[nums[i]] = cnt + 1
		} else {
			set[nums[i]] = 1
		}
		if set[nums[i]] > res {
			res = set[nums[i]]
		}
	}
	return res
}

func main() {
	nums := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	fmt.Println(longestConsecutive1(nums))
	fmt.Println(longestConsecutive2(nums))
}
