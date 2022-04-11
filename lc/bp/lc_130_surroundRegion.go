package main

import "fmt"

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
找到所有与边缘O相连的字母O
*/
//bfs
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

// dfs
func solve130(board [][]byte) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || i == m-1 || j == 0 || j == n-1
			if isEdge && board[i][j] == 'O' {
				dfs130(i, j, board)
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

func dfs130(i, j int, board [][]byte) {
	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '$' || board[i][j] == 'X' {
		return
	}
	board[i][j] = '$'
	dfs130(i+1, j, board)
	dfs130(i-1, j, board)
	dfs130(i, j-1, board)
	dfs130(i, j+1, board)
}

//并查集

func solve130II(board [][]byte) {
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
						connect(i*n+j, i*n+j-1, arr)
					}
					if board[i][j+1] == 'O' {
						connect(i*n+j, i*n+j+1, arr)
					}
				}
			}
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(arr[i], " ")
		if i > 0 && (i+1)%n == 0 {
			fmt.Println()
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && find(i*n+j, arr) != find(m*n, arr) {
				board[i][j] = 'X'
			}
		}
	}
}

func setFa(arr []int) {
	for i := range arr {
		arr[i] = i
	}
}

func find(i int, arr []int) int {
	root := i
	for root != arr[root] {
		root = arr[root]
	}
	for i != arr[i] {
		i, arr[i] = arr[i], root
	}
	return root
}

func connect(i, j int, arr []int) {
	ii, jj := find(i, arr), find(j, arr)
	if ii != jj {
		arr[ii] = jj
	}
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != root {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
	}
}

func (u *UnionFind) IsConnected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}

func solveIII(board [][]byte) {
	if len(board) <= 2 {
		return
	}

	var (
		m, n         = len(board), len(board[0])
		dummy, ufind = m * n, NewUnionFind(m*n + 1)
		position     func(i, j int) int
	)
	position = func(i, j int) int {
		return i*n + j
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i == 0 || j == 0 || i == m-1 || j == n-1 {
					ufind.Union(position(i, j), dummy)
				}
				if i > 0 && board[i-1][j] == 'O' {
					ufind.Union(position(i, j), position(i-1, j))
				}
				if j > 0 && board[i][j-1] == 'O' {
					ufind.Union(position(i, j), position(i, j-1))
				}
			}
		}
	}
	for i := 0; i < len(ufind.parent)-1; i++ {
		fmt.Print(ufind.parent[i], " ")
		if i > 0 && (i+1)%n == 0 {
			fmt.Println()
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !ufind.IsConnected(position(i, j), dummy) {
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

	board2 := [][]byte{
		{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'O', 'O', 'X', 'X', 'X', 'X', 'O'},
		{'X', 'X', 'X', 'X', 'O', 'O', 'O', 'O'},
		{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'X'},
		{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'X', 'X', 'O', 'O', 'X', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'O', 'X', 'X'}}

	board2 = [][]byte{{'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'},
		{'O', 'X', 'O', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'},
		{'O', 'X', 'O', 'O', 'O', 'O'}}

	show(board2)
	fmt.Println()
	solve130II(board2)
	//solveIII(board2)
	show(board2)
}
