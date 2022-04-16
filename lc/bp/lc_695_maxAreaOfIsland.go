package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
岛屿最大面积
给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。
岛屿的面积是岛上值为 1 的单元格的数目。
计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
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
	res := 1
	m, n := len(grid), len(grid[0])
	queue := make([]int, 0)
	queue = append(queue, i*n+j)
	dirX := []int{-1, 0, 1, 0}
	dirY := []int{0, -1, 0, 1}
	for len(queue) > 0 {
		front := queue[0]
		x, y := front/n, front%n
		queue = queue[1:]

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

// 并查集
var max int

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent, size: size}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for root != u.parent[root] {
		root = u.parent[root]
	}
	for u.parent[i] != root {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u *UnionFind) Union(x, y int) {
	xx, yy := u.Find(x), u.Find(y)
	if xx != yy {
		u.parent[xx] = yy
		u.size[yy] += u.size[xx]
		max = util.MaxInt(max, u.size[yy])
	}

}

func maxAreaOfIslandIII(grid [][]int) int {
	var (
		m, n     = len(grid), len(grid[0])
		ufind    = NewUnionFind(m*n + 1)
		position func(i, j int) int
	)
	position = func(i, j int) int {
		return i*n + j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				k := position(i, j)
				ufind.parent[k] = k
				ufind.size[k] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i+1 < m && grid[i+1][j] == 1 {
					ufind.Union(position(i+1, j), position(i, j))
				}
				if j+1 < n && grid[i][j+1] == 1 {
					ufind.Union(position(i, j+1), position(i, j))
				}
			}
		}
	}
	return max
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
	fmt.Println(maxAreaOfIsland(grid))
	grid = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIslandII(grid))

	grid = [][]int{
		{1}}
	fmt.Println(maxAreaOfIslandIII(grid))
}
