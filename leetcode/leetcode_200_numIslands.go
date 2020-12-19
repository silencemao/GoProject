package main

import "fmt"

/*
寻找孤岛的点的个数
*/

func numIslands(grid [][]byte) int {
	cnt := 0
	if len(grid) < 1 {
		return cnt
	}
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfsNum(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func dfsNum(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 0
	dfsNum(grid, i+1, j)
	dfsNum(grid, i-1, j)
	dfsNum(grid, i, j+1)
	dfsNum(grid, i, j-1)
}

func numIslandsBFS(grid [][]byte) int {
	cnt := 0
	if len(grid) < 1 || len(grid[0]) < 1 {
		return cnt
	}
	m, n := len(grid), len(grid[0])
	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	var queue []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			grid[i][j] = '0'
			cnt++
			queue = append(queue, i*n+j)
			for len(queue) > 0 {
				front := queue[0]
				queue = queue[1:]
				for k := 0; k < 4; k++ {
					x, y := front/n+dirX[k], front%n+dirY[k]
					if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
						continue
					}
					grid[x][y] = '0'
					queue = append(queue, x*n+y)
				}
			}
		}
	}
	return cnt
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
	fmt.Println(numIslandsBFS(grid))
}
