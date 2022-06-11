package main

import _struct "GoProject/leetcode/struct"

/*
单词搜索进阶版 类似第79题
给一个单词列表，一个二维矩阵，找出二维矩阵中可以组成的单词
*/

func inArea212(x, y int, board [][]byte) bool {
	m, n := len(board), len(board[0])
	if x >= 0 && x < m && y >= 0 && y < n {
		return true
	}
	return false
}

func help212(board [][]byte, visited [][]bool, word string, x, y, k int) bool {
	if board[x][y] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}

	visited[x][y] = true
	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, -1, 0, 1}
	result := false
	for i := 0; i < 4; i++ {
		newX, newY := x+dirX[i], y+dirY[i]
		if inArea212(newX, newY, board) && !visited[newX][newY] && help212(board, visited, word, newX, newY, k+1) {
			result = true
			break
		}
	}
	visited[x][y] = false
	return result
}

func exist212(board [][]byte, word string) bool {
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
			if help212(board, visited, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

/*
超时
*/
func findWords(board [][]byte, words []string) []string {
	var res []string
	for _, word := range words {
		if exist212(board, word) {
			res = append(res, word)
		}
	}
	return res
}

/*
前缀树
*/
func findWords1(board [][]byte, words []string) []string {
	res := []string{}
	m := len(board)
	if m == 0 || len(words) == 0 {
		return res
	}
	n := len(board[0])
	if n == 0 {
		return res
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	// 1. 把所有单词构造成前缀树
	trie := _struct.Constructor()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	track := []byte{}
	// 2. 对矩阵每一个点进行回溯，找出所有在前缀树中的单词，注意找到了之后要修改trie对应节点的 isEnd
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(board, visited, i, j, &trie, &track, &res)
		}
	}

	return res
}

func backtrack(board [][]byte, visited [][]bool, i, j int, parent *_struct.Trie, track *[]byte, res *[]string) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] {
		return
	}
	c := board[i][j]
	// trie 中不存在该字符
	if parent.Next[c-'a'] == nil {
		return
	}
	// trie 存在就可以设置访问标志，添加路径
	visited[i][j] = true
	*track = append(*track, c)
	// trie node 切换到该字符分支上，如果 isEnd 为true可以插入结果
	child := parent.Next[c-'a']
	if child.IsEnd {
		*res = append(*res, string(*track))
		child.IsEnd = false // 防止重复插入
	}
	// 4个方向DFS
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, v := range directions {
		backtrack(board, visited, i+v[0], j+v[1], child, track, res)
	}
	// 回溯恢复访问标志和路径
	visited[i][j] = false
	*track = (*track)[:len(*track)-1]
}

func main() {

}
