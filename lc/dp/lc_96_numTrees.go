package main

import "fmt"

/*
节点数为n的二叉搜索树的形状个数
*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	//dp[2] = 2
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	n := 3
	fmt.Println(numTrees(n))
}
