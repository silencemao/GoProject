package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
元素的 频数 是该元素在一个数组中出现的次数。
给你一个整数数组 nums 和一个整数 k 。在一步操作中，你可以选择 nums 的一个下标，并将该下标对应元素的值增加 1 。
执行最多 k 次操作后，返回数组中最高频元素的 最大可能频数 。

输入：nums = [1,2,4], k = 5
输出：3
解释：对第一个元素执行 3 次递增操作，对第二个元素执 2 次递增操作，此时 nums = [4,4,4] 。
4 是数组中最高频元素，频数是 3

*/

// 排序+滑动窗口
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	res := 1
	l, r := 0, 1
	total := int64(0)
	for ; r < len(nums); r++ {
		total += int64((nums[r] - nums[r-1]) * (r - l))
		for total > int64(k) {
			total -= int64(nums[r] - nums[l])
			l += 1
		}
		res = util.MaxInt(res, r-l+1)
	}
	return res
}

func maxFrequency1(nums []int, k int) int {
	sort.Ints(nums)
	sums := make([]int, len(nums)+1)
	res := 1
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	//k = k % sums[len(sums)-1]
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(sums)-1
		for l < r {
			mid := l + (r-l)>>1
			tmp := nums[mid]*(mid-i+1) - (sums[mid+1] - sums[i])
			if tmp > k {
				r = mid
			} else {
				l = mid + 1
			}
		}
		res = util.MaxInt(res, l-i)
	}

	return res
}

func main() {
	nums := []int{1, 2, 4, 8, 13}
	k := 119
	fmt.Println(maxFrequency(nums, k))
	fmt.Println(maxFrequency1(nums, k))
}
