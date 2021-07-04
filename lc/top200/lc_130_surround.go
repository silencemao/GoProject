package main

import "fmt"

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
func help130(board [][]byte, i, j int) {
	m, n := len(board), len(board[0])
	if 0 <= i && i < m && 0 <= j && j < n {
		if board[i][j] == 'O' {
			board[i][j] = '1'

			if i > 0 && board[i-1][j] == 'O' {
				help130(board, i-1, j)
			}
			if i+1 < m && board[i+1][j] == 'O' {
				help130(board, i+1, j)
			}
			if j > 0 && board[i][j-1] == 'O' {
				help130(board, i, j-1)
			}
			if j+1 < m && board[i][j+1] == 'O' {
				help130(board, i, j+1)
			}
		}
	}
}

func solve(board [][]byte) {
	if len(board) < 1 || len(board[0]) < 1 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && (i == 0 || i == m-1) || (j == 0 || j == n-1) {
				help130(board, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

/*
上一个的做法是沿着边界找，找到边界为'O'以及与边界的'O'联通的'O'

改成从内向外找，找到一个'O' 然后向外扩展，看四个方向联通的部分 若存在边界的'O' 则不改变当前的值
若四个方向不存在与边界'O'相连，则改变当前位置的值
*/

func help1302(board [][]byte, i, j int, visited *[][]bool) bool {
	m, n := len(board), len(board[0])
	if i < 0 || i >= m-1 || j < 0 || j >= n-1 {
		return false
	}
	if (*visited)[i][j] == true || board[i][j] == 'X' {
		return true
	}
	(*visited)[i][j] = true
	flag := help1302(board, i-1, j, visited) && help1302(board, i+1, j, visited) && help1302(board, i, j-1, visited) && help1302(board, i, j+1, visited)
	(*visited)[i][j] = false
	return flag
}

func solve1(board [][]byte) {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && help1302(board, i, j, &visited) {
				board[i][j] = '1'
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

/*
bfs求解，非递归（队列）
*/

func solve2(board [][]byte) {
	queue := make([]int, 0)
	dirX := []int{-1, 0, 1, 0}
	dirY := []int{0, -1, 0, 1}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				queue = append(queue, i*n+j)
			} else {
				continue
			}
			board[i][j] = '1'
			for len(queue) > 0 {
				front := queue[0]
				queue = queue[1:]
				for k := 0; k < len(dirX); k++ {
					newX, newY := front/n+dirX[k], front%n+dirY[k]
					if 0 <= newX && newX < m && 0 <= newY && newY < n && board[newX][newY] == 'O' {
						board[newX][newY] = '1'
						queue = append(queue, newX*n+newY)
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'O', 'X', 'O'},
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
	}
	solve2(board)
	fmt.Println(board)
}
