package main

import "fmt"

/*
更新矩阵
给定二维矩阵，由0和1组成 若当前位置为0，则不做处理，若当前位置为1 则计算距离其最近的0的距离，只计算上下左右四个方向
*/
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return [][]int{}
	}
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, m)
	var queue [][]int
	for i := range res {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = matrix[i][j]
			if res[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				res[i][j] = -1
			}
		}
	}

	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for i := 0; i < len(dirY); i++ {
			x, y := front[0]+dirX[i], front[1]+dirY[i]
			if x < 0 || x >= m || y < 0 || y >= n || res[x][y] != -1 {
				continue
			}
			queue = append(queue, []int{x, y})
			res[x][y] = res[front[0]][front[1]] + 1
		}

	}

	return res
}

func main() {
	matrix := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(matrix))
}
