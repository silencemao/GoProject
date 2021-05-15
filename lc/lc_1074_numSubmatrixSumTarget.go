package main

import "fmt"

func numSubMatrixSumTarget(matrix [][]int, target int) int {
	res := 0

	sumMatrix := make([][]int, len(matrix)+1)
	for i := range sumMatrix {
		sumMatrix[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(sumMatrix); i++ {
		for j := 1; j < len(sumMatrix[i]); j++ {
			sumMatrix[i][j] = sumMatrix[i][j-1] + sumMatrix[i-1][j] - sumMatrix[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	for i := 1; i < len(sumMatrix); i++ {
		for j := 1; j < len(sumMatrix[0]); j++ {
			for p := 1; p <= i; p++ {
				for q := 1; q <= j; q++ {
					t := sumMatrix[i][j] - sumMatrix[i][q-1] - sumMatrix[p-1][j] + sumMatrix[p-1][q-1]
					if t == target {
						res++
					}
				}
			}
		}
	}

	return res
}

func numSubMatrixSumTarget1(matrix [][]int, target int) int {
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		for j := i; j < len(matrix[0]); j++ {
			cnt := 0
			sumDict := make(map[int]int) // 思考为什么这个dict不是全局的，因为不连续？
			sumDict[0] = 1
			for k := 0; k < len(matrix); k++ {
				if i > 0 {
					cnt += matrix[k][j] - matrix[k][i-1]
				} else {
					cnt += matrix[k][j]
				}
				if cnt > target {
					if _, ok := sumDict[cnt-target]; ok {
						res += sumDict[cnt-target]
					}
				}
				if cnt == target {
					res += sumDict[cnt-target]
				}
				sumDict[cnt] += 1
			}
		}
	}

	return res
}

func main() {
	matrix := [][]int{{0, 1, 0, 0, 1}, {0, 0, 1, 1, 1}, {1, 1, 1, 0, 1}, {1, 1, 0, 1, 1}, {0, 1, 1, 0, 0}}
	target := 1
	fmt.Println(numSubMatrixSumTarget(matrix, target))
	fmt.Println(numSubMatrixSumTarget1(matrix, target))
	a := make(map[int]int)
	fmt.Println(a)
	a[10] += 1
	fmt.Println(a)
}
