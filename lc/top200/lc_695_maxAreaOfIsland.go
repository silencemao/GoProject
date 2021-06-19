package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
同200
*/
func islandArea(grid [][]int, i, j int, area *int) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		if grid[i][j] == 1 {
			grid[i][j] = 2
			*area += 1
			islandArea(grid, i-1, j, area)
			islandArea(grid, i+1, j, area)
			islandArea(grid, i, j-1, area)
			islandArea(grid, i, j+1, area)
		}
	}
}

func islandArea1(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		if grid[i][j] == 1 {
			grid[i][j] = 2
			return 1 + islandArea1(grid, i-1, j) +
				islandArea1(grid, i+1, j) +
				islandArea1(grid, i, j-1) +
				islandArea1(grid, i, j+1)
		}
	}
	return 0
}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				//area := 0
				//islandArea(grid, i, j, &area)
				area := islandArea1(grid, i, j)
				res = util.MaxInt(res, area)
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
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
