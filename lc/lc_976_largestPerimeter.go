package main

import (
	"fmt"
	"sort"
)

/*

给定一个数组，计算数组中元素可以组成三角形的最大周长
*/

func largestPerimeter(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	sort.Ints(arr)
	for i := len(arr) - 1; i >= 2; i-- {
		if arr[i] < arr[i-1]+arr[i-2] {
			return arr[i] + arr[i-1] + arr[i-2]
		}
	}
	return 0
}

func main() {
	arr := []int{3, 2, 1}
	fmt.Println(largestPerimeter(arr))
}
