package main

import "fmt"

/*
四数相加，给定4个数组，找出有多少个组合可以使得A[i]+B[j]+C[k]+c[m]=0
两两hash表，前两个数组求和存储到hash表，另外两个数组求和sum，判断是否在hash表中存在-sum
A+B+C=T转换为A+B、T-C是否存在问题
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	set1 := make(map[int]int, 0)
	res := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			if cnt, ok := set1[sum]; ok {
				set1[sum] = cnt + 1
			} else {
				set1[sum] = 1
			}
		}
	}
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			sum := nums3[i] + nums4[j]
			if cnt, ok := set1[-sum]; ok {
				res += cnt
			}
		}
	}
	return res
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
