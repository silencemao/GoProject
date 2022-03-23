package main

import "fmt"

/*
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	//for i := 2; i <= n; i++ { // dp[n-i]无法递推
	//	dp[i] = dp[i-1] * dp[n-i]
	//}
	for i := 2; i <= n; i++ { // 节点个数
		for j := 1; j <= i; j++ { // 以j为根节点
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	n := 5
	fmt.Println(numTrees(n))
}
