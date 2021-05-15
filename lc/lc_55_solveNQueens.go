package main

import (
	"fmt"
	"math"
)


func isSafe(numLevel int, flag *[]int) bool {
	for i := 1; i < numLevel; i++ {
		if ((*flag)[numLevel] == (*flag)[i]) || (int64(math.Abs(float64(numLevel-i))) == int64(math.Abs(float64((*flag)[numLevel]-(*flag)[i])))) {
			return false
		}
	}
	return true
}
func queenHelper(n int, numLevel int, flag *[]int) {
	if numLevel > n {
		res := make([][]string, n)
		for i := 0; i < n; i++ {
			res[i] = make([]string, n)
			for j := 0; j < n; j++ {
				res[i][j] = "."
			}
		}
		for i := 1; i <= n; i++ {
			(res)[i-1][(*flag)[i] - 1] = "Q"
			fmt.Print((*flag)[i] - 1, " ")
		}
		fmt.Println()
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print((res)[i][j], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	} else {
		for i := 1; i <= n; i++ {
			(*flag)[numLevel] = i
			if isSafe(numLevel, flag) {
				queenHelper(n, numLevel+1, flag)
			}
		}
	}
}

func solveNQueens(n int) [][]string {
	res := make([][]string, n)
	for i := 0; i < n; i++ {
		res[i] = make([]string, n)
	}
	flag := make([]int, n+1)
	queenHelper(n, 1, &flag)
	return res
}

func main() {
	solveNQueens(8)
}
