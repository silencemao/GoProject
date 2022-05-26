package main

import (
	"GoProject/leetcode/util"
	"container/list"
	"fmt"
	"sort"
)

func check(x, y, n int) bool {
	if x >= 0 && x < n && y >= 0 && y < n {
		return true
	}
	return false
}

func swimInWater(grid [][]int) int {
	res := grid[0][0]
	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, -1, 0, 1}
	m := len(grid)
	var queue []int
	queue = append(queue, 0)
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		x, y := front/m, front%m
		if x == m-1 && y == m-1 {
			break
		}
		grid[x][y] = -1
		tmp, useX, useY := 1<<31-1, 0, 0
		for k := 0; k < 4; k++ {
			newX, newY := x+dirX[k], y+dirY[k]
			if check(newX, newY, m) && grid[newX][newY] != -1 && tmp >= grid[newX][newY] {
				tmp = grid[newX][newY]
				useX, useY = newX, newY
				grid[newX][newY] = -1
				queue = append(queue, useX*m+useY)
			}
		}
		if useX+useY > 0 {
			//queue = append(queue, useX*m+useY)
		}
		if tmp != 1<<31-1 {
			res = util.MaxInt(res, tmp)
		}
	}
	return res
}

func canReach(h int, grid [][]int) bool {
	const maxN = 50
	n := len(grid)
	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}
	visited := [maxN][maxN]bool{} // 标识已经访问过的平台，也可以用map；其实应该是n*n大小，但是用n的话代码要多几行，需要遍历初始化每一行
	visited[0][0] = true
	set := list.New()
	set.PushBack([]int{0, 0}) // 用长度为2的切片代表一个点；初始位置放入集合
	for set.Len() > 0 {
		pos := set.Remove(set.Back()).([]int)
		x, y := pos[0], pos[1]
		if x == n-1 && y == n-1 {
			return true
		}
		for i := 0; i < len(dr); i++ {
			newX, newY := x+dr[i], y+dc[i]
			if newX >= 0 && newX < n && newY >= 0 && newY < n &&
				!visited[newX][newY] && grid[newX][newY] <= h && grid[x][y] <= h { // grid[x][y] <= h这个也很关键，h必须大于出发点才行对于{{3,2},{1,0}}
				set.PushBack([]int{newX, newY})
				visited[newX][newY] = true
			}
		}
	}
	return false
}

func swimInWater1(grid [][]int) int {
	n := len(grid)
	l, r := 0, n*n
	for l < r {
		mid := (l + r) >> 1
		if canReach(mid, grid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater2(grid [][]int) (ans int) {
	n := len(grid)
	return sort.Search(n*n-1, func(threshold int) bool {
		if threshold < grid[0][0] {
			return false
		}
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < n && !vis[x][y] && grid[x][y] <= threshold {
					vis[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return vis[n-1][n-1]
	})
}

func main() {
	grid := [][]int{{3, 2},
		{0, 1}}
	//fmt.Println(swimInWater(grid))
	fmt.Println(swimInWater1(grid))
	fmt.Println(swimInWater2(grid))
}
