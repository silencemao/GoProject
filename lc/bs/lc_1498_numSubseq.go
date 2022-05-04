package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 nums 和一个整数 target 。
请你统计并返回 nums 中能满足其最小元素与最大元素的 和 小于或等于 target 的 非空 子序列的数目。
由于答案可能很大，请将结果对109+ 7取余后返回。


1、当nums[l]+nums[r]<=target时，l+1 ~ r之间的任意一个数字作为右边界都满足此条件
2、n个数字的组合个数共有多少个，C(n,0) + C(n,1)+..+C(n,n) = 2^n
*/
// 双指针
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] * 2 % (1e9 + 7) // 表示长度为i时，组合个数，比如长度2的数组，组合个数为4
	}
	l, r, res := 0, len(nums)-1, 0
	for l < r {
		tmp := nums[l] + nums[r]
		if tmp <= target {
			res = (res + dp[r-l]) % (1e9 + 7) //表示以l为左边界，l+1至r之间任一数字作为右边界，都满足上述条件，共计组合个数
			l += 1
		} else {
			r -= 1
		}
	}
	return res
}

/*
二分法
*/
func numSubseq1(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 0
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] * 2 % (1e9 + 7) // 表示长度为i时，组合个数，比如长度2的数组，组合个数为4
	}
	for i := 0; i < len(nums) && nums[i]*2 <= target; i++ { // 确定左边界，二分法查找右边界的最大索引
		max := target - nums[i]
		l, r := i+1, len(nums)-1
		for l <= r {
			mid := l + (r-l)>>1
			if nums[mid] <= max {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		//if l >= len(nums) || nums[l] > max {
		//	l -= 1
		//}
		if r-i >= 0 {
			res = (res + dp[r-i]) % (1e9 + 7)
		}
	}
	return res
}

func main() {
	a := (1e9) + 7

	fmt.Println(a)

	nums := []int{3, 5, 6, 7}
	target := 9
	fmt.Println(numSubseq(nums, target))
	fmt.Println(numSubseq1(nums, target))
}
