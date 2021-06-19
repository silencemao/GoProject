package main

import "fmt"

/*
岛屿的周长

同200、695
*/

func perimeterHelp(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i >= 0 && i < m && j >= 0 && j < n {
		if grid[i][j] == 1 {
			grid[i][j] = 2
			return perimeterHelp(grid, i-1, j) + perimeterHelp(grid, i+1, j) + perimeterHelp(grid, i, j-1) + perimeterHelp(grid, i, j+1)
		} else if grid[i][j] == 0 { // 岛屿遇到海 边界加一
			return 1
		}
	} else { // 岛屿走出边界，边界加一
		return 1
	}
	return 0
}

func islandPerimeter(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += perimeterHelp(grid, i, j)
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{{1, 0, 0}}
	fmt.Println(islandPerimeter(grid))
}
