package main

import "fmt"

func generate(numsRows int) [][]int {
	var res [][]int
	if numsRows < 1 {
		return res
	}
	res = append(res, []int{1})
	if numsRows == 1 {
		return res
	}
	res = append(res, []int{1, 1})
	if numsRows == 2 {
		return res
	}
	for i := 2; i < numsRows; i++ {
		tmp := make([]int, i+1)
		for j := 1; j <= i-1; j++ {
			tmp[j] = res[i-1][j-1] + res[i-1][j]
		}
		tmp[0], tmp[i] = 1, 1
		res = append(res, tmp)
	}
	fmt.Println(res)
	return res
}

func main() {
	fmt.Println(generate(5))
}
