package main

import (
	"fmt"
	"math"
)

//有序数组找出距离数字x最近的k个数字
/*
给定一个排序好的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
*/
func findClosestElements(arr []int, k, x int) []int {
	for len(arr) > k {
		l, r := 0, len(arr)-1
		if math.Abs(float64(arr[l]-x)) <= math.Abs(float64(arr[r]-x)) {
			arr = arr[:r]
		} else {
			arr = arr[1:]
		}
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k, x := 4, 3
	fmt.Println(findClosestElements(arr, k, x))
}
