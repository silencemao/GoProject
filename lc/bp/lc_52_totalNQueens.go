package main

import (
	"fmt"
)

/*
n皇后问题 研究的是如何将 n个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

*/
func help52(res *int, row, n int, curRow *[][]string) {
	if row == n {
		*res += 1
		return
	}
	for col := 0; col < n; col++ {
		if isValid52(row, col, n, *curRow) {
			(*curRow)[row][col] = "Q"
			help52(res, row+1, n, curRow)
			(*curRow)[row][col] = "."
		}
	}
}

func isValid52(row, col, n int, curRow [][]string) bool {
	//检查列
	for i := 0; i < row; i++ {
		if curRow[i][col] == "Q" {
			return false
		}
	}
	//检查45度角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if curRow[i][j] == "Q" {
			return false
		}
	}

	//检查135度角
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if curRow[i][j] == "Q" {
			return false
		}
	}

	return true
}

func totalNQueens(n int) int {
	var res int
	tmp := make([][]string, n)
	for i := range tmp {
		tmp[i] = make([]string, n)
		for j := range tmp[i] {
			tmp[i][j] = "."
		}
	}

	help52(&res, 0, n, &tmp)
	return res
}

func main() {
	n := 4
	fmt.Println(totalNQueens(n))
}
