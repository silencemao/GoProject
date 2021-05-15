package main

import (
	"fmt"
	"GoProject/leetcode/util"
)

/*
给定一个字符串数组，每一个字符串均是由 0 1组成的
m表示给定的0的个数 n表示给定的1的个数，利用现有的m个0， n个1可以组成几个数组中的字符串

Input: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
Output: 4

5个0 3个1可以组成"10"、"0001"、"1"、"0"，返回4

可以采用动态规划的思路，dp[i][j]表示i个0 j个1可以组成的字符串个数，
遍历每一个字符串，统计其中的0和1的个数，num0,num1。dp[i-num0][j-num1] 在在拼成str之前，可以拼成的字符串个数。
然后dp[i][j] = util.MaxInt(dp[i - num0][j - num1] + 1, dp[i][j])，即可
*/

func findMaxForm(strs []string, m int, n int) int {
	if len(strs) < 1{
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		b := []byte(str)
		num1, num0 := 0, 0
		for _, char := range b {
			if char=='1' {
				num1++
			} else {
				num0++
			}
		}

		for i := m; i >= num0; i-- {
			for j := n; j >= num1; j-- {
				dp[i][j] = util.MaxInt(dp[i - num0][j - num1] + 1, dp[i][j])
			}
		}
	}
	return dp[m][n]
}

func main() {
	strs := []string{"10", "0", "1"}
	m, n := 1, 1
	fmt.Println(findMaxForm(strs, m, n))
}
