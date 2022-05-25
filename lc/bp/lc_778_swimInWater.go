package main

import (
	"GoProject/leetcode/util"
	"fmt"
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

func main() {
	grid := [][]int{{10, 12, 4, 6},
		{9, 11, 3, 5},
		{1, 7, 13, 8},
		{2, 0, 15, 14}}
	fmt.Println(swimInWater(grid))
}
