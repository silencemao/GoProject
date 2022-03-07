package main

/*
解数独
*/
func solveSudoku(board [][]byte) {
	help37(board)
}

func isValid37(board [][]byte, row, col int, k byte) bool {
	m, n := len(board), len(board[0])
	//行检查
	for i := 0; i < n; i++ {
		if board[row][i] == k {
			return false
		}
	}
	//列检查
	for i := 0; i < m; i++ {
		if board[i][col] == k {
			return false
		}
	}
	//九宫格检查
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := startRow; i < startRow+3; i += 1 {
		for j := startCol; j < startCol+3; j += 1 {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}

func help37(board [][]byte) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := '1'; k <= '9'; k++ {
				if isValid37(board, i, j, byte(k)) {
					board[i][j] = byte(k)
					if help37(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func main() {

}
