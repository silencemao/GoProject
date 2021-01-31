package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个数组，找出其中的排列，保证相邻两个数之和为平方数，返回有多少个这样的排列
全排列的另一种考察方法，数组中可能会存在重复，也就是说本题是去重全排列的一种考察方法
判断两数之和是否为平方数 pow(int(sqrt(a+b)), 2) == a+b
*/
func numSquarefulPerms(A []int) int {
	var res int
	sort.Ints(A)
	helper996(A, 0, &res)
	return res
}

func helper996(A []int, start int, res *int) {
	if start >= len(A) {
		*res++
	}

	m := make(map[int]bool)
	for i := start; i < len(A); i++ {
		if _, ok := m[A[i]]; ok {
			continue
		}
		m[A[i]] = true
		A[i], A[start] = A[start], A[i]
		if start == 0 || (start > 0 && isSquare(float64(A[start]+A[start-1]))) {
			helper996(A, start+1, res)
		}
		A[i], A[start] = A[start], A[i]
		m[A[i]] = false
	}
}

func isSquare(num float64) bool {
	return math.Pow(float64(int64(math.Sqrt(num))), 2) == num
}

func main() {
	A := []int{80, 1, 80, 1, 3, 6, 3}
	fmt.Println(numSquarefulPerms(A))
}
