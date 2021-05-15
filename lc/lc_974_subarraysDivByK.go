package main

import (
	"fmt"
)

/*
给定一个数组arr，计算arr中有多少个连续的子数组的和可以整除k

*/
func subarraysDivByK(arr []int, k int) int {
	res := 0
	sum := 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range arr {
		sum = (num%k + k + sum) % k
		if cnt, ok := m[sum]; ok {
			res += cnt
			m[sum] += 1
		} else {
			m[sum] = 1
		}
	}

	return res
}

func subarraysDivByK2(arr []int, k int) int {
	res, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1

	for i, _ := range arr {
		sum = (sum + arr[i]) % k
		if sum < 0 {
			sum += k
		}
		if cnt, ok := m[sum]; ok {
			res += cnt
			m[sum] += 1
		} else {
			m[sum] = 1
		}
	}
	return res
}

// 超时
func subarraysDivByK1(arr []int, k int) int {
	res := 0
	if len(arr) == 1 && arr[0]%k == 0 {
		return 1
	}
	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		if sum%k == 0 {
			res++
		}
		for j := i - 1; j >= 0; j-- {
			sum += arr[j]
			if sum%k == 0 {
				res++
			}
		}
	}
	return res
}

func main() {
	arr := []int{4, 5, 0, -2, -3, 1}
	k := 5
	fmt.Println(subarraysDivByK(arr, k))
	fmt.Println(subarraysDivByK2(arr, k))
	fmt.Println(subarraysDivByK1(arr, k))
}
