package main

import (
	"fmt"
	"sort"
)
/*
给定一个数组，以及一个数字K
选择i位置上的数字，让A[i] = -A[i]
i位置可以重复选择，返回K次操作之后数组和的最大值
*/
func largestSumAfterKNegations(A []int, K int) int {
	res := 0

	for i := 0; i < K; i++ {
		sort.Ints(A)
		A[0] = -A[0]
	}
	for i := 0; i < len(A); i++ {
		res += A[i]
	}

	return res
}

func main() {
	A := []int{2,-3,-1,5,-4}
	K := 2
	fmt.Println(largestSumAfterKNegations(A, K))
}
