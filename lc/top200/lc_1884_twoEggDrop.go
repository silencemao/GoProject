package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func help1884(k, n int, dp *[][]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	if (*dp)[n][k] != 0 {
		return (*dp)[n][k]
	}
	res := 1<<31 - 1
	for i := 1; i <= n; i++ {
		res = util.MinInt(res, util.MaxInt(help1884(k, n-i, dp), help1884(k-1, i-1, dp))+1)
	}
	(*dp)[n][k] = res
	return res
}

func twoEggDrop(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {

		dp[i] = make([]int, 3)
	}
	fmt.Println(help1884(2, n, &dp))
	//fmt.Println(dp)
	return dp[n][2]
}

func main() {
	n := 100
	fmt.Println(twoEggDrop(n))
}
