package main

import "fmt"

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
*/
func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	dirX := []int{-1, 1, 0, 0}
	dirY := []int{0, 0, -1, 1}
	var queue []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				queue = append(queue, i*n+j)
				board[i][j] = '$'
			} else {
				continue
			}

			for len(queue) > 0 {
				front := queue[0]
				queue = queue[1:]
				for k := 0; k < 4; k++ {
					newX, newY := front/n+dirX[k], front%n+dirY[k]
					if newX < 0 || newX >= m || newY < 0 || newY >= m || board[newX][newY] != 'O' {
						continue
					}
					board[newX][newY] = '$'
					queue = append(queue, newX*n+newY)
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '$' {
				board[i][j] = 'O'
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
	solve(board2)
	show(board2)
}
