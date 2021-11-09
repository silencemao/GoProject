package main

import "fmt"
/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

n = 3
[
	[1 2 3]
	[8 9 4]
	[7 6 5]
]
顺时针添加数字，左闭右闭，然后更新行/列
*/
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	target := n*n

	num := 1
	top, bottom := 0, n-1
	left, right := 0, n-1
	for num <= target {
		//从左到右
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		//从上到下
		for i := top; i<=bottom; i++ {
			res[i][right] = num
			num++
		}
		right--
		//从右到左
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--
		//从下到上
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++

	}
	return res
}
func main() {
	n := 3
	res := generateMatrix(n)
	for i := range res {
		fmt.Println(res[i])
	}
}
