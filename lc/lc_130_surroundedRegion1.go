package main

import (
	"fmt"
)

func setFa(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
}

func find(x int, arr []int) int {
	if x != arr[x] {
		arr[x] = find(arr[x], arr)
	}
	return arr[x]
}

func connect(i, j int, arr []int) {
	ii := find(i, arr)
	jj := find(j, arr)
	if ii < jj {
		arr[i] = jj
	}
	if ii > jj {
		arr[j] = ii
	}
}

func solve2(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	arr := make([]int, m*n+1)
	setFa(arr)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					connect(i*n+j, m*n, arr)
				} else {
					if board[i-1][j] == 'O' {
						connect(i*n+j, (i-1)*n+j, arr)
					}
					if board[i+1][j] == 'O' {
						connect(i*n+j, (i+1)*n+j, arr)
					}
					if board[i][j-1] == 'O' {
						connect(i*n+j, i*n+(j-1), arr)
					}
					if board[i][j+1] == 'O' {
						connect(i*n+j, i*n+(j+1), arr)
					}
				}
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
		if i > 0 && (i+1)%n == 0 {
			fmt.Println()
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if find(i*n+j, arr) != n*m {
				board[i][j] = 'X'
			}
		}
	}
}

func show(arr [][]byte) {
	m, n := len(arr), len(arr[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%q ", arr[i][j])
		}
		fmt.Println()
	}
}

func main() {

	//
	/*board := [][]byte{
	{'X', 'O', 'X', 'X'},
	{'O', 'X', 'O', 'X'},
	{'X', 'O', 'X', 'O'},
	{'O', 'X', 'O', 'X'},
	{'X', 'O', 'X', 'O'},
	{'O', 'X', 'O', 'X'}} */
	/*board := [][]byte{
	{'X', 'X', 'O', 'X'},
	{'X', 'O', 'O', 'X'},
	{'X', 'O', 'O', 'X'},
	{'X', 'X', 'O', 'X'}} */

	//board1 := [][]byte{
	//	{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'O'},
	//	{'O', 'O', 'O', 'X', 'X', 'X', 'X', 'O'},
	//	{'X', 'X', 'X', 'X', 'O', 'O', 'O', 'O'},
	//	{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'X'},
	//	{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
	//	{'O', 'X', 'X', 'O', 'O', 'X', 'X', 'O'},
	//	{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
	//	{'O', 'X', 'X', 'X', 'X', 'O', 'X', 'X'}}
	board2 := [][]byte{
		{'X', 'X', 'X', 'O'},
		{'X', 'O', 'O', 'O'},
		{'X', 'O', 'O', 'O'},
		{'X', 'X', 'X', 'X'},
	}
	show(board2)
	fmt.Println()
	solve2(board2)
	show(board2)
}
