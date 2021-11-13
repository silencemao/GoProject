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

func fourSumCount1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) {
	set := make(map[int][]int, 0)
	for i1, v1 := range nums1 {
		for i2, v2 := range nums2 {
			sum := v1 + v2
			if tmp, ok := set[sum]; ok {
				tmp[0] += 1
				tmp = append(tmp, []int{i1, i2}...)
				//set[sum] = tmp
			} else {
				tmp = append(tmp, 1)
				tmp = append(tmp, []int{i1, i2}...)
				set[sum] = tmp
			}
		}
	}

	res := make([][]int, 0)
	for i3, v3 := range nums3 {
		for i4, v4 := range nums4 {
			sum := v3 + v4
			if tmp, ok := set[-sum]; ok {
				for i := 0; i < tmp[0]; i++ {
					res = append(res, []int{tmp[i*2], tmp[i*2+1], i3, i4})
				}
			}
		}
	}
	fmt.Println(res)

}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
	fourSumCount1(A, B, C, D)
}
