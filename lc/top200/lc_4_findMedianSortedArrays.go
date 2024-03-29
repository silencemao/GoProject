package main

import (
	"fmt"
)

/*
寻找两个排序数组的中位数
O(m+n)，若两个数组的长度之和为偶数个，则返回索引为(m+n)/2 - 1， (m+n)/2, 两个位置数的平均数
        若两个数组的长度之和为奇数个，则返回索引为(m+n)/2的数
所以无论是奇数还是偶数，都需要遍历(m+n)/2 + 1次
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	prev, cur, i, j := 0, 0, 0, 0
	for k := 0; k <= length/2; k++ {
		prev = cur
		// i >= m时，说明nums1遍历完了，从nums2中开始寻找，若j>=n 从nums1中开始寻找
		if i < m && (j >= n || nums1[i] < nums2[j]) {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}
	}
	if length&1 == 0 {
		return float64(cur+prev) / 2
	} else {
		return float64(cur)
	}
}

func find1(nums1, nums2 []int) float64 {
	prev, cur := 0, 0
	m, n := len(nums1), len(nums2)
	k := m + n
	i, j, k := 0, 0, 0
	for ; k <= (m+n)/2; k++ {
		prev = cur
		if i < m && (nums1[i] < nums2[j] || j >= n) {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}
	}
	if (m+n)&1 == 0 {
		return float64(prev+cur) / 2.0
	}
	return float64(cur)
}

func f4(nums1, nums2 []int) float64 {
	l, r := len(nums1), len(nums2)
	prev, cur := 0, 0
	ind1, ind2 := 0, 0
	for i := 0; i <= (l+r)/2; i++ {
		if ind1 < l && ind2 < r {
			if nums1[ind1] <= nums2[ind2] {
				prev, cur = cur, nums1[ind1]
				ind1 += 1
			} else {
				prev, cur = cur, nums2[ind2]
				ind2 += 1
			}
		} else if ind1 < l {
			prev, cur = cur, nums1[ind1]
			ind1 += 1
		} else {
			prev, cur = cur, nums2[ind2]
			ind2 += 1
		}

	}
	if (l+r)&1 == 1 {
		return float64(cur)
	}
	return float64(prev+cur) / 2.0
}

func main() {
	nums1 := []int{1, 3, 4, 5, 7, 8}
	nums2 := []int{2, 4, 6, 7}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(find1(nums1, nums2))
	fmt.Println(f4(nums1, nums2))

}
