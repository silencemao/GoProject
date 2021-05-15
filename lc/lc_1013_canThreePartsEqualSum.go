package main

import (
	"fmt"
)
/*
给定一个数组，判断该数组是否可以拆分成三个子数组，每个子数组的和相等
若既然能将一个数组拆分成三个
*/
func canThreePartsEqualSum(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	if sum % 3 != 0 {
		return false
	}
	partSum := 0
	cnt, i := 0, 0
	for i = range arr {
		partSum += arr[i]
		if partSum == sum / 3 {
			partSum = 0
			cnt++
			if cnt == 3 {
				break
			}
		}
	}
	if cnt == 3 && i == len(arr) - 1 {
		return true
	}
	return false
}

func main() {
	nums := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	fmt.Println(canThreePartsEqualSum(nums))
}
