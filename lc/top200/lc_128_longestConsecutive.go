package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
最好O(n)时间解决
*/
/*
排序 + 动态规划 看i位置的值与i-1位置的值的关系
1、相等 则dp[i]=dp[i-1]
2、nums[i]=nums[i-1]+1 dp[i]=d[i-1]+1
3、其它 dp[i]=1
*/

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			dp[i] = dp[i-1] + 1
		} else if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = 1
		}
		res = util.MaxInt(dp[i], res)
	}
	return res
}

func longestConsecutive1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := 1
	set := make(map[int]int)
	sort.Ints(nums)
	set[nums[0]] = 1
	for i := range nums {
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
	nums := []int{100, 4, 200, 2, 3, 1}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive1(nums))
}
