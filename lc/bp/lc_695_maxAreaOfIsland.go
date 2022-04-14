package main

import "fmt"

/*
岛屿最大面积
*/
func dfs695(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
		return 0
	}
	if grid[i][j] == 1 {
		grid[i][j] = 2
		return 1 + dfs695(grid, i-1, j) + dfs695(grid, i+1, j) + dfs695(grid, i, j-1) + dfs695(grid, i, j+1)
	}
	return 0
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cur := dfs695(grid, i, j)
				if cur > res {
					res = cur
				}
			}
		}
	}
	return res
}

func bfs695(grid [][]int, i, j int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	queue := make([]int, 0)
	queue = append(queue, i*n+j)
	dirX := []int{-1, 0, 1, 0}
	dirY := []int{0, -1, 0, 1}
	for len(queue) > 0 {
		front := queue[0]
		x, y := front/n, front%n
		queue = queue[1:]

		res += 1
		grid[i][j] = 2
		for k := 0; k < 4; k++ {
			nexX, nexY := x+dirX[k], y+dirY[k]
			if nexX >= 0 && nexX < m && nexY >= 0 && nexY < n && grid[nexX][nexY] == 1 {
				grid[nexX][nexY] = 2
				res += 1
				queue = append(queue, nexX*n+nexY)
			}
		}
	}
	return res
}

func maxAreaOfIslandII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cur := bfs695(grid, i, j)
				if cur > res {
					res = cur
				}
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	//fmt.Println(maxAreaOfIsland(grid))
	fmt.Println(maxAreaOfIslandII(grid))
}
