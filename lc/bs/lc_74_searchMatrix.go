package main

import "fmt"

/*
编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
*/

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := 0
	for ; i < m; i++ {
		if matrix[i][n-1] >= target {
			break
		}
	}
	if i >= m {
		return false
	}
	for j := 0; j < n; j++ {
		if matrix[i][j] == target {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
