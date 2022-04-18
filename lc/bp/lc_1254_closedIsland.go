package main

import "fmt"

/*
封闭岛屿的数目
*/
func bfs1254(grid [][]int, i, j int) int {
	m, n, res := len(grid), len(grid[0]), 1
	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, -1, 0, 1}
	queue := []int{i*n + j}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		x, y := front/n, front%n
		grid[i][j] = 2
		for k := 0; k < 4; k++ {
			newX, newY := x+dirX[k], y+dirY[k]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] == 1 {
				continue
			}
			if grid[newX][newY] == 0 {
				grid[newX][newY] = 2
				queue = append(queue, newX*n+newY)
			}
			if grid[newX][newY] == 0 && (newX == 0 || newX == m-1 || newY == 0 || newY == n-1) {
				res = 0
			}
		}
	}
	return res
}

func closedIsland(grid [][]int) int {
	m, n, res := len(grid), len(grid[0]), 0

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				res += bfs1254(grid, i, j)
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{{0, 0, 1, 1, 0, 1, 0, 0, 1, 0}, {1, 1, 0, 1, 1, 0, 1, 1, 1, 0}, {1, 0, 1, 1, 1, 0, 0, 1, 1, 0}, {0, 1, 1, 0, 0, 0, 0, 1, 0, 1}, {0, 0, 0, 0, 0, 0, 1, 1, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 1, 1, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 1, 1, 0, 1, 1, 0, 1, 1, 0}}
	fmt.Println(grid)
	fmt.Println(closedIsland(grid))
}
