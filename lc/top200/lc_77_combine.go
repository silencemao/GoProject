package main

import "fmt"

/*
给定整数 n, k，返回从1-n中任意k个数的组合，任一组合内数据无重复
*/
func combine77(n int, k int) [][]int {
	var res [][]int
	var tmp []int
	help77(&res, tmp, 1, n, k)
	for _, r := range res {
		fmt.Println(r)
	}
	return res
}

func help77(res *[][]int, tmp []int, ind, n, k int) {
	if len(tmp) == k {
		*res = append(*res, append([]int{}, tmp...))
		return
	}
	//for i := ind; i <= n; i++ {
	for i := ind; i <= n-(k-len(tmp))+1; i++ { // 剪枝，i的最大值处进行处理
		tmp = append(tmp, i)
		help77(res, tmp, i+1, n, k) // 此处使用ind+1的话会出现重复
		tmp = tmp[:len(tmp)-1]
	}
}

func main() {
	combine77(4, 2)
}
