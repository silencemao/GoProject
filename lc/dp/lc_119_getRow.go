package main

/*杨辉三角
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/
func getRow(rowIndex int) []int {
	C := make([][]int, rowIndex+1)
	for i := range C {
		C[i] = make([]int, i+1)
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
	return C[rowIndex]
}

func main() {

}
