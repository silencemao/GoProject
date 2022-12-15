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

var max int

type UnionFind2 struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind2 {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind2{parent: parent, size: size}
}

func (u *UnionFind2) Find(i int) int {
	root := i
	for root != u.parent[root] {
		root = u.parent[root]
	}
	for u.parent[i] != root {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u *UnionFind2) Union(x, y int) {
	xx, yy := u.Find(x), u.Find(y)
	if xx != yy {
		u.parent[xx] = yy
		u.size[yy] += u.size[xx]
		max = MaxInt(max, u.size[yy])
	}

}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxAreaOfIsland2(grid [][]int) int {
	var (
		m, n     = len(grid), len(grid[0])
		ufind    = NewUnionFind(m*n + 1)
		position func(i, j int) int
	)
	position = func(i, j int) int {
		return i*n + j
	}
	max = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				k := position(i, j)
				ufind.parent[k] = k
				ufind.size[k] = 1
				if max != 1 {
					max = 1
				}
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
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	//fmt.Println(maxAreaOfIsland(grid))
	fmt.Println(maxAreaOfIsland2(grid))
}
