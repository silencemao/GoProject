package main

import (
	"fmt"
	"math"
	"sort"
)

// 查找第一个缺失的正整数
func firstMissingPositive(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	sort.Ints(nums)
	if nums[0] > 1 {
		return 1
	} else {
		pre := -1
		for i := 0; i < len(nums); i++ {
			if nums[i] > 0 {
				if pre == -1 {
					if nums[i] > 1 {
						return 1
					} else {
						pre = nums[i]
					}
				}
				if nums[i]-pre > 1 {
					return pre + 1
				} else {
					pre = nums[i]
				}
			}
		}
	}
	if nums[len(nums)-1]+1 <= 0 {
		return 1
	}
	return nums[len(nums)-1] + 1
}

func firstMissingPositive1(nums1 []int) int {
	sort.Ints(nums1)
	min := 1
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == min {
			min++
		}
	}
	return min
}

func firstMissingPositive2(nums2 []int) int {
	n := len(nums2)
	for i := 0; i < n; i++ { // 先将<=0的数 变为 n+1
		if nums2[i] <= 0 {
			nums2[i] = n + 1
		}
	}
	for _, num := range nums2 { // 将1<=abs(num) <=n的数，所对应的索引位置变为负数，表示当前位置的数在[1,n]的范围内
		num = int(math.Abs(float64(num)))
		if num <= n {
			nums2[num-1] = -int(math.Abs(float64(nums2[num-1]))) // 此处要注意
		}
	}
	for i, num := range nums2 { // i所对应的索引位置 > 0，表示当前位置的数不在[1,n]的范围内，是缺失值
		if num > 0 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	nums := []int{1, 1}
	fmt.Println(firstMissingPositive2(nums))
}
