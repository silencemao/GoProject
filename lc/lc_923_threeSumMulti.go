package main

import (
	"fmt"
	"sort"
)
/*

Given an integer array A, and an integer target, return the number of tuples i, j, k  such that i < j < k and A[i] + A[j] + A[k] == target.

As the answer can be very large, return it modulo 10^9 + 7.
三数之和的多种组合
给定一个数组A和一个目标值target
找出满足和为target并且索引值是从小到大排序的组合个数
*/
func threeSumMulti(A []int, target int) int {
	res := 0
	sort.Ints(A)
	for i := 0; i < len(A) - 2; i++ {
		tmp := target - A[i]
		left, right := i + 1, len(A) - 1
		for left < right {
			if A[left] + A[right] > tmp {
				right--
			} else if A[left] + A[right] < tmp {
				left++
			} else {
				lCount, rCount := 1, 1
				for left + lCount < right && A[left + lCount]==A[left] {
					lCount++
				}
				for left + lCount <= right - rCount && A[right - rCount]==A[right] {
					rCount++
				}
				if A[left]==A[right] {
					res += (right - left + 1) * (right - left) / 2 % (1e9 + 7)
				} else {
					res += lCount * rCount % (1e9 + 7)
				}
				left += lCount
				right -= rCount
			}
		}
	}
	return res
}

func threeSumMulti1(A []int, target int) int {
    res := 0
	tmpMap := make(map[int]int)
	for _, num := range A {
		tmpMap[num] += 1
	}
	for num1, cnt1 := range tmpMap {
		for num2, cnt2 := range tmpMap {
			subNum := target - num1 - num2
			if cnt3, ok := tmpMap[subNum]; ok {
				if num1 == num2 && num2 == subNum {
					res += cnt1 * (cnt1 - 1) * (cnt1 - 2) / 6 % (1e9 + 7)
				} else if num1 == num2 && num2 != subNum {
					res += cnt1 * (cnt2 - 1) * cnt3 / 2 % (1e9 + 7)
				} else if num1 < num2 && num2 < subNum {
					res += cnt1 * cnt2 * cnt3 % (1e9 + 7)
				}
			}
		}
	}
	return res % (1e9+7)
}

func main() {
	A := []int{1,1,2,2, 3, 3, 4, 4, 5, 5}
	target := 8
	fmt.Println(threeSumMulti1(A, target))
}