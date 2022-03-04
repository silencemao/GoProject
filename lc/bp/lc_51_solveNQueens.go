package main

import (
	"fmt"
	"strings"
)

/*
N皇后
*/
func help51(res *[][]string, row, n int, curRow *[][]string) {
	if row == n {
		var tmp []string
		for _, r := range *curRow {
			tmp = append(tmp, strings.Join(r, ""))
		}
		*res = append(*res, tmp)
		return
	}
	for col := 0; col < n; col++ {
		if isValid(row, col, n, *curRow) {
			(*curRow)[row][col] = "Q"
			help51(res, row+1, n, curRow)
			(*curRow)[row][col] = "."
		}
	}
}

func isValid(row, col, n int, curRow [][]string) bool {
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

func solveNQueens(n int) [][]string {
	var res [][]string
	tmp := make([][]string, n)
	for i := range tmp {
		tmp[i] = make([]string, n)
		for j := range tmp[i] {
			tmp[i][j] = "."
		}
	}

	help51(&res, 0, n, &tmp)
	return res
}

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}
