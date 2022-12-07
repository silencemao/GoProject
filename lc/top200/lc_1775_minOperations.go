package main

import (
	"fmt"
	"sort"
)

/*
给你两个长度可能不等的整数数组nums1 和nums2。两个数组中的所有值都在1到6之间（包含1和6）。
每次操作中，你可以选择 任意数组中的任意一个整数，将它变成 1到 6之间 任意的值（包含 1和 6）。
请你返回使 nums1中所有数的和与nums2中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1。

假设 nums1元素求和 < nums2元素求和
d = sum(nums2) - sum(nums1)
nums1中的每个元素最多可以增加 6-num, nums2中的元素最多可以减少 num-1，将这些增量/差值 存到数组arr中，逆序排序
从 d 中减去arr中的元素，当d<= 0时，则存在最少操作次数
*/
func sum(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}

func minOperations1775(nums1 []int, nums2 []int) int {
	s1, s2 := sum(nums1), sum(nums2)
	if s1 == s2 {
		return 0
	}
	if s1 > s2 {
		return minOperations1775(nums2, nums1)
	}
	d := s2 - s1
	var arr []int
	for _, num := range nums1 {
		arr = append(arr, 6-num)
	}
	for _, num := range nums2 {
		arr = append(arr, num-1)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for i, v := range arr {
		d -= v
		if d <= 0 {
			return i + 1
		}
	}
	return -1
}

func minOperations17751(nums1 []int, nums2 []int) int {
	s1, s2 := sum(nums1), sum(nums2)
	if s1 == s2 {
		return 0
	}
	if s1 > s2 {
		return minOperations17751(nums2, nums1)
	}
	d := s2 - s1
	cnt := [6]int{}
	for _, num := range nums1 {
		cnt[6-num] += 1
	}
	for _, num := range nums2 {
		cnt[num-1] += 1
	}
	ans := 0
	for i := 5; i >= 0; i-- {
		for cnt[i] > 0 && d > 0 {
			d -= i
			cnt[i] -= 1
			ans += 1
			if d <= 0 {
				return ans
			}
		}
	}
	return -1
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{1, 1, 2, 2, 2, 2}
	fmt.Println(minOperations1775(nums1, nums2))
	fmt.Println(minOperations17751(nums1, nums2))
}
