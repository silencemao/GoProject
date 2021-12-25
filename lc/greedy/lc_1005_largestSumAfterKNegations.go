package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

选择某个下标 i并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。

*/
func largestSumAfterKNegations(nums []int, k int) int {
	res := 0
	sort.Ints(nums)
	if nums[0] >= 0 { // 全为正数
		if k%2 == 1 {
			nums[0] = -nums[0]
		}
	} else if nums[len(nums)-1] <= 0 { // 全为负数
		i := 0
		for ; i < k && i < len(nums); i++ {
			nums[i] = -nums[i]
		}
		if i < k && (k-i)%2 == 1 {
			nums[len(nums)-1] = -nums[len(nums)-1]
		}
	} else {
		negCnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= 0 {
				negCnt++
			}
		}

		if k >= negCnt { // k大于负数个数, 负数取反
			for i := 0; i < negCnt && nums[i] <= 0; i++ {
				nums[i] = -nums[i]
			}
			sort.Ints(nums)
			if (k-negCnt)%2 == 1 {
				nums[0] = -nums[0]
			}
		} else { // k小于负数个数
			for i := 0; i < k && i < negCnt; i++ {
				nums[i] = -nums[i]
			}
		}
	}
	for _, v := range nums {
		res += v
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

/*
按照绝对值从大到小排序，
遍历数组，只要k>0就把数组中的负数取反
剩余k对绝对值最小的数字取反
*/
func largestSumAfterKNegations1(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func largestSumAfterKNegations2(nums []int, k int) int {
	sort.Ints(nums)
	for i := 0; i < k; i++ {
		if i < len(nums) && nums[i] < 0 {
			nums[i] = -nums[i]
		} else {
			sort.Ints(nums)
			if (k-i)%2 == 1 {
				nums[0] = -nums[0]
			}
			break
		}
	}
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func main() {
	nums := []int{2, -3, -1, 5, -4}
	k := 2
	fmt.Println(largestSumAfterKNegations(nums, k))
}
