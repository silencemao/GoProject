package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func help474(str string) (int, int) {
	z, o := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i]-'0' == 0 {
			z += 1
		} else {
			o += 1
		}
	}
	return z, o
}

/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n ，strs中的字符串由0和1组成
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 ，
返回包含m个0 n个1的子集的最大长度

二维dp数组，dp[i][j]表示包含i个0和j个1的子集的长度，
dp[0][0] = 1具有0个0和0个1的最大子集长度为0
dp[i][j] = max(dp[i][j], dp[i-z][j-o]+1)
*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0
	for _, str := range strs {
		z, o := help474(str)
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = util.MaxInt(dp[i][j], dp[i-z][j-o]+1)
			}
		}
	}

	return dp[m][n]
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	fmt.Println(findMaxForm(strs, m, n))
}
