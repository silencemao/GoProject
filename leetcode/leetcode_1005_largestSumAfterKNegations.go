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
	sort.Ints(A)
	for i := 0; i < K; i++ {
		if i < len(A) && A[i] <= 0 {
			A[i] = -A[i]
		} else { // 全部变为正数之后，根据剩余次数的奇偶分别做处理，偶数不做处理，奇数对最小的取反
			sort.Ints(A)
			if (K-i)%2 != 0 {
				A[0] = -A[0]
			}
			break
		}
	}
	for i := 0; i < len(A); i++ {
		res += A[i]
	}
	return res
}

func largestSumAfterKNegations1(A []int, K int) int {
	res := 0
	sort.Ints(A)
	for i := 0; K > 0 && i < len(A) && A[i] <= 0; i, K = i+1, K-1 { // A[i] <= 0的判断一定要在for循环中
		A[i] = -A[i]
	}
	min := 1<<32 - 1
	for i := range A {
		res += A[i]
		if A[i] < min {
			min = A[i]
		}
	}
	if K%2 != 0 {
		res = res - min*2
	}
	return res
}

func main() {
	A := []int{4, 2, 3}
	K := 1
	//fmt.Println(largestSumAfterKNegations(A, K))
	fmt.Println(largestSumAfterKNegations1(A, K))
}
