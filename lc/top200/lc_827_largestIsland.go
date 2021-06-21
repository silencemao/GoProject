package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格0 变成1 。
返回执行此操作后，grid 中最大的岛屿面积是多少？
岛屿 由一组上、下、左、右四个方向相连的1 形成。

1、使用ind 标记联通的岛屿
2、areas记录每一个岛屿的面积areas[ind]=area
3、遍历grid[i][j]=0的点找其上下左右的 四个位置的ind（此处用set记录）
4、对ind对应位置的面积相加1+areas[ind]+...

https://leetcode-cn.com/problems/making-a-large-island
*/
func calArea(grid [][]int, i, j, ind int) int {
	m, n := len(grid), len(grid[0])
	if i >= 0 && i < m && j >= 0 && j < n {
		if grid[i][j] == 1 {
			grid[i][j] = ind
			return 1 + calArea(grid, i-1, j, ind) + calArea(grid, i+1, j, ind) + calArea(grid, i, j-1, ind) + calArea(grid, i, j+1, ind)
		}
	}
	return 0
}

func fillOcean(grid [][]int, i, j int, areas []int) int {
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	m, n := len(grid), len(grid[0])
	set := make(map[int]bool)
	for k := 0; k < 4; k++ {
		nr, nc := i+dr[k], j+dc[k]
		if nr >= 0 && nr < m && nc >= 0 && nc < n {
			if grid[nr][nc] > 1 {
				set[grid[nr][nc]] = true
			}
		}
	}
	area := 1
	for k := range set {
		area += areas[k]
	}
	return area
}

func largestIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	ind := 2
	areas := make([]int, m*n+2)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := calArea(grid, i, j, ind)
				res = util.MaxInt(res, area)
				areas[ind] = area
				ind += 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				area := fillOcean(grid, i, j, areas)
				res = util.MaxInt(res, area)
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{{1, 0}, {1, 0}}
	fmt.Println(largestIsland(grid))
}
