package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	if m==0 || n==0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 先处理第一行和第一列
	for i := 0; i < m; i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			if i > 0 {
				dp[i][0] = dp[i-1][0]
			}
		}
	}

	for j := 0; j < n; j++ {
		if text1[0] == text2[j] {
			dp[0][j] = 1
		} else {
			if j > 0 {
				dp[0][j] = dp[0][j-1]
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i]==text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = util.MaxInt(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m-1][n-1]
}

/*
给定两个字符串 text1 text2,返回两个字符串的最长公共子串（不要求子串连续）
       a  b  c  d  e
    0  0  0  0  0  0
  a 0  1  1  1  1  1
  c 0  1  1  2  2  2
  e 0  1  1  2  2  3

以text1="abcde" text2="ace"为例，
1、首先生成一个(m+1, n+1)大小的矩阵，m+1 n+1的目的是当两个字符串为空的时候依然可以满足，同时为了后面比较时方便
2、text1[i]==text2[j]时，即text1以i结尾和text2以j结尾时，二者的最长公共子序列的长度如何计算
   看上图横纵坐标均为a a的时候，text1 text2均以a为结尾的时候，最长公共子序列长度为1，因为他们两个a前面均是空，公共长度为0，
   当同时向后移动一步到a时，公共长度加1 dp[i][j] = dp[i-1][j-1] + 1
3、text1[i]!=text2[j]时，即text1以i结尾和text2以j结尾时
   看上图横坐标为c，纵坐标为d时，二者的公共子序列为2
   既然以i和j结尾时，字符不相等，那就要看以i和j-1为结尾时的公共长度  以及 以i-1和j结尾时的公共长度
   dp[i][j] = max(dp[i-1][j], dp[i][j-1])

所以本题的转移方程是 2 和 3中的两个公式，分别对应text1[i]==text2[j] 和 text1[i] != text2[j]的时候
*/
func longestCommonSubsequence1(text1, text2 string) int {
	m, n := len(text1), len(text2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1]==text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = util.MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func main() {
	text1, text2 := "bsbininm", "jmjkbkjkv"
	fmt.Println(longestCommonSubsequence1(text1, text2))
}
