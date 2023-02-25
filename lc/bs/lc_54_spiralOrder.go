package main

import "fmt"

/*
螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/
func spiralOrder(matrix [][]int) []int {
	var res []int
	l, r, t, d := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t += 1
		if t > d {
			break
		}

		for i := t; i <= d; i++ {
			res = append(res, matrix[i][r])
		}
		r -= 1
		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			res = append(res, matrix[d][i])
		}
		d -= 1
		if d < t {
			break
		}

		for i := d; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l += 1
		if l > r {
			break
		}
	}
	return res
}

// 逆时针打印
func spiralOrder1(matrix [][]int) []int {
	var res []int
	l, r, t, d := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := t; i <= d; i++ {
			res = append(res, matrix[i][l])
		}
		l += 1
		if l > r {
			break
		}

		for i := l; i <= r; i++ {
			res = append(res, matrix[d][i])
		}
		d -= 1
		if d < t {
			break
		}

		for i := d; i >= t; i-- {
			res = append(res, matrix[i][r])
		}
		r -= 1
		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			res = append(res, matrix[t][i])
		}
		t += 1
		if t > d {
			break
		}
	}
	return res
}

func main() {
	matrix := [][]int{{2, 5, 8},
		{4, 0, 1},
		{3, 2, 1}}
	fmt.Println(spiralOrder(matrix))
	fmt.Println(spiralOrder1(matrix))
}
