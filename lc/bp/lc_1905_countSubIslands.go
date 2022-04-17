package main

import "fmt"

/*
给你两个m x n的二进制矩阵grid1 和grid2，它们只包含0（表示水域）和 1（表示陆地）。一个 岛屿是由 四个方向（水平或者竖直）上相邻的1组成的区域。
任何矩阵以外的区域都视为水域。

如果 grid2的一个岛屿，被 grid1的一个岛屿完全 包含，也就是说 grid2中该岛屿的每一个格子都被 grid1中同一个岛屿完全包含，那么我们称 grid2中的这个岛屿为 子岛屿。
请你返回 grid2中 子岛屿的 数目。

两个网格矩阵，若grid2中的岛屿能完全被grid1中的岛屿覆盖，则岛屿数量加1。
完全覆盖指grid2中标记为1的位置，grid2中也是为1，若grid2中岛屿的某一个位置 在grid1中为0，则不是完全覆盖

1、遍历grid2，若grid2[i][j] = 1 && grid1[i][j] = 1，则从i，j位置处开始bfs搜索
2、将i，j位置记录入队列，并将grid2[i][j] = 0，防止bfs中重复遍历
3、遍历grid2中i，j的四个方向，若newX,newY位置处为1，则加入队列，并设置为0
4、判断grid1[newX][newY]是否为1，若不为1，则res=0，此时先不要返回，继续把queue中的元素以及周边的元素遍历完，保证i，j可覆盖的范围都标记为0，
防止一半为0，一半不为0，出现误判
*/
func bfs1905(grid1, grid2 [][]int, i, j int) int {
	dirX := []int{-1, 0, 1, 0}
	dirY := []int{0, -1, 0, 1}
	m, n := len(grid2), len(grid2[0])
	queue := []int{i*n + j}
	res := grid1[i][j]
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		x, y := front/n, front%n
		grid2[x][y] = 0
		for k := 0; k < 4; k++ {
			newX, newY := x+dirX[k], y+dirY[k]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && grid2[newX][newY] == 1 {
				grid2[newX][newY] = 0
				queue = append(queue, newX*n+newY)
				if grid1[newX][newY] != 1 {
					res = 0
				}
			}
		}
	}
	return res
}
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	m, n := len(grid1), len(grid1[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res += bfs1905(grid1, grid2, i, j)
			}
		}
	}
	return res
}

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	for _, a := range grid1 {
		fmt.Println(a)
	}
	fmt.Println()
	for _, a := range grid2 {
		fmt.Println(a)
	}
	fmt.Println(countSubIslands(grid1, grid2))
	//for _, a := range grid1 {
	//	fmt.Println(a)
	//}

}
