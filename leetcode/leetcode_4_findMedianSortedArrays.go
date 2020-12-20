package main

import (
	"fmt"
	"math"
)

/*
寻找两个排序数组的中位数 时间复杂度控制在O(log(m+n)) m n 为两个数组的长度。看到这个时间复杂度，应该可以想到二分法了
1、中位数的位置就是就是使得它左右两边的数字个数相等，如果我们知道第一个数组应该在何处切分，那就可以推算出第二个数组切分的位置
例如：m=4 n=6，第一个数组切分的位置是2，则第二个数组切分的位置就是3，所以我们就可以对较短的数组进行切分，第一个数组的切分位置确定之后，
第二个数组切分的位置自然就确定了。
2、首先定义第一个数组切分位置左右两边的数字分别为L1 R1 同理第二个数组为L2 R2
切分完之后若L1<=R2 && L2<=R1，则切分位置就是正确的，不需要调整切分位置了
若L1 > R2 说明较短数组切分之后右边数字太大了，需要左移切分位置
若L2 > R1 说明较短数组切分之后左边数字太小了，需要右移切分位置
*/
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	if len(nums1) > len(nums2) {
//		nums1, nums2 = nums2, nums1
//	}
//	if len(nums1) == 0 {
//		if len(nums2)%2 == 0 {
//			return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2-1]) / 2.0
//		} else {
//			return float64(nums2[len(nums2)/2])
//		}
//	}
//	length := len(nums1) + len(nums2)
//	cut1, cut2 := 0, 0 // 切分之后第一、二数组右边的起始位置
//	L1, L2, R1, R2 := 0, 0, 0, 0
//	cutL, cutR := 0, length // 左右边界
//	for cut1 <= len(nums1) {
//		cut1 = (cutL + cutR) / 2
//		if cut1 >= len(nums1) {
//			cut1 = len(nums1)
//		}
//		cut2 = length/2 - cut1
//		if cut1 == 0 { // 第一个全部切分成右边 L1 > R2 永远不成立
//			L1 = math.MinInt64
//		} else {
//			L1 = nums1[cut1-1]
//		}
//
//		if cut1 == len(nums1) { // 第一个数组全部被切成左边 L2 > R1永远不成立
//			R1 = math.MaxInt64
//		} else {
//			R1 = nums1[cut1]
//		}
//
//		if cut2 == 0 {
//			L2 = math.MinInt64
//		} else {
//			L2 = nums2[cut2-1]
//		}
//		if cut2 == len(nums2) {
//			R2 = math.MaxInt64
//		} else {
//			R2 = nums2[cut2]
//		}
//
//		if L1 > R2 {
//			cutR = cut1 - 1
//		} else if L2 > R1 {
//			cutL = cut1
//		} else {
//			if length%2 == 1 {
//				return math.Min(float64(R1), float64(R2))
//			}
//
//			if length%2 == 0 {
//				return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2
//			}
//		}
//	}
//	return -1
//}

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
func findMedianSortedArrays1(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2-1]) / 2.0
		} else {
			return float64(nums2[len(nums2)/2])
		}
	}
	m, n := len(nums1), len(nums2)

	left, right := 0, m
	totalLeft := (m + n + 1) / 2

	for left < right { // 确定位置i 使nums1[i-1] <= nums2[totalLeft-i]
		i := (left + right + 1) / 2 //
		j := totalLeft - i
		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}
	i, j := left, totalLeft-left
	nums1LeftMax, nums2LeftMax := math.MinInt64, math.MinInt64
	nums1RightMin, nums2RightMin := math.MaxInt64, math.MaxInt64

	if i != 0 {
		nums1LeftMax = nums1[i-1]
	}
	if i != m {
		nums1RightMin = nums1[i]
	}

	if j != 0 {
		nums2LeftMax = nums2[j-1]
	}
	if j != n {
		nums2RightMin = nums2[j]
	}
	if (m+n)%2 == 0 {
		return (math.Max(float64(nums1LeftMax), float64(nums2LeftMax)) + math.Min(float64(nums1RightMin), float64(nums2RightMin))) / 2
	} else {
		return math.Max(float64(nums1LeftMax), float64(nums2LeftMax))
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 7}
	fmt.Println(findMedianSortedArrays1(nums1, nums2))

}
