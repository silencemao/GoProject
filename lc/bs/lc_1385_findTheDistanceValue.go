package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
给你两个整数数组arr1，arr2和一个整数d，请你返回两个数组之间的距离值。
「距离值」定义为符合此距离要求的元素数目：对于元素arr1[i]，不存在任何元素arr2[j]满足 |arr1[i]-arr2[j]| <= d 。
*/

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	cnt := 0
	sort.Ints(arr1)
	sort.Ints(arr2)
	m, n := len(arr1), len(arr2)
	for i := 0; i < m; i++ {
		j := 0
		for ; j < n; j++ {
			if util.Abs(arr1[i]-arr2[j]) <= d {
				break
			}
		}
		if j >= n {
			cnt += 1
		}
	}
	return cnt
}

func main() {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2
	fmt.Println(findTheDistanceValue(arr1, arr2, d))
}
