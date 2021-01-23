package main

import (
	"fmt"
)

/*
给定一个排序之后的数组，数组有正数和负数，返回数组中元素平方之后 从小到大的顺序
*/
func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	i, j := 0, len(A)-1
	for k := len(A) - 1; k >= 0; k-- {
		if abs(A[i]) > abs(A[j]) {
			res[k] = A[i] * A[i]
			i++
		} else {
			res[k] = A[j] * A[j]
			j--
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func main() {
	arr := []int{-6, -5, 0, 2, 5}
	fmt.Println(sortedSquares(arr))
}
