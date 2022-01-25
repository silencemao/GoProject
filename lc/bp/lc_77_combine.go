package main

import "fmt"

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。
*/
func help77(res *[][]int, tmp []int, n, k, ind int) {
	if len(tmp) == k {
		*res = append(*res, append([]int{}, tmp...))
		return // TODO 此处是否一定需要return
	}
	for i := ind; i <= n-(k-len(tmp))+1; i++ {
		//for i := ind; i <= n; i++ {
		tmp = append(tmp, i)
		help77(res, tmp, n, k, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func combine(n int, k int) [][]int {
	var res [][]int
	help77(&res, []int{}, n, 2, 1)
	//fmt.Println(res)
	return res
}
func main() {
	fmt.Println(combine(4, 2))
}
