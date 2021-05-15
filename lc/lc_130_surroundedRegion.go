package main

/*
给定一个二维矩阵，里面填充的是'X' 或者 'O'，将被'X'包围的'O'转换为'X'
若'O'处于边界或者与边界的'O'相连，则不进行转换

本题的本质就是将四周全部被'X'包围的'O'转换为'X'，所以边界出的'O'是不需要被转换的，
以及与边界处的'O'相连的'O'也不需要被转换。
因此识别边界处的'O'以及与边界相连的'O'是本题的关键

一、DFS识别
遍历四个边界，若点为'O'，则进行DFS搜索，将与其相连的'O'都识别出来
二、BFS搜索

*/
/*
1、扫描边缘的'O'，遇到'O'则，将与其相连的'O'变成'$'
2、将所有剩下的'O'变成'X'
3、将'$'变成'O'
*/
func solve(board [][]byte) {
	if len(board) < 1 || len(board[0]) < 1 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				solveDFS(board, i, j)
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

func solveDFS(board [][]byte, i, j int) {
	m, n := len(board), len(board[0])
	if board[i][j] == 'O' {
		board[i][j] = '$'
		if i > 0 && board[i-1][j] == 'O' {
			solveDFS(board, i-1, j)
		}
		if i < m-1 && board[i+1][j] == 'O' {
			solveDFS(board, i+1, j)
		}
		if j > 0 && board[i][j-1] == 'O' {
			solveDFS(board, i, j-1)
		}
		if j < n-1 && board[i][j+1] == 'O' {
			solveDFS(board, i, j+1)
		}
	}
}

/*
BFS
*/

func solve1(board [][]byte) {
	if len(board) < 1 || len(board[0]) < 1 {
		return
	}
	m, n := len(board), len(board[0])
	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	var queue []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				queue = append(queue, i*n+j)
			} else {
				continue
			}
			board[i][j] = '$'
			for len(queue) > 0 {
				front := queue[0]
				queue = queue[1:]
				for k := 0; k < 4; k++ {
					x, y := front/n+dirX[k], front%n+dirY[k]
					if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
						continue
					}
					board[x][y] = '$'
					queue = append(queue, x*n+y)
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

func main() {

}
