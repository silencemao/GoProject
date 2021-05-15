package main

import "fmt"

func inArea(board [][]byte, i, j int) bool {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
		return true
	}
	return false
}

func existHelp(board [][]byte, visited [][]bool, word string, i, j, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}

	visited[i][j] = true
	result := false
	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range direction {
		newi, newj := i+dir[0], j+dir[1]
		if inArea(board, newi, newj) && !visited[newi][newj] {
			res := existHelp(board, visited, word, newi, newj, k+1)
			if res {
				result = true
				break
			}
		}
	}
	visited[i][j] = false
	return result
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return true
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if existHelp(board, visited, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
