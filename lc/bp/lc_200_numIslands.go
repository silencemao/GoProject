package main

import "fmt"

/*
统计岛屿数量
*/
func bfs200(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	queue := []int{i*n + j}
	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, -1, 0, 1}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		x, y := front/n, front%n
		grid[x][y] = '0'
		for k := 0; k < 4; k++ {
			newX, newY := x+dirX[k], y+dirY[k]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] == '0' {
				continue
			}
			if grid[newX][newY] == '1' {
				grid[newX][newY] = '0' // 这一步是防止重复把某一个位置加入到队列中，内存暴增
				queue = append(queue, newX*n+newY)
			}
		}
	}
}
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res += 1
				bfs200(grid, i, j)
			}
		}
	}
	return res
}

func dfs200(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs200(grid, i+1, j)
		dfs200(grid, i-1, j)
		dfs200(grid, i, j+1)
		dfs200(grid, i, j-1)
	}
}

func numIslands1(grid [][]byte) int {
	m, n, res := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res += 1
				dfs200(grid, i, j)
			}
		}
	}
	return res
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
	grid = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands1(grid))
}
