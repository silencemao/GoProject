package main

import (
	"fmt"
)

/*
给定三个字符串s1、s2、s3，请你帮忙验证s3是否是由s1和s2 交错 组成的。

两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
注意：a + b 意味着字符串 a 和 b 连接。
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, k := len(s1), len(s2), len(s3)
	if m+n != k {
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[m][n]
}

func isInterleave1(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return dfs97(s1, s2, s3, 0, 0, 0)
}

func dfs97(s1, s2, s3 string, k1, k2, k3 int) bool {
	if k1+k2 >= len(s3) {
		return true
	}
	if k1 < len(s1) && s1[k1] == s3[k3] {
		if dfs97(s1, s2, s3, k1+1, k2, k3+1) {
			return true
		}

	}
	if k2 < len(s2) && s2[k2] == s3[k3] {
		if dfs97(s1, s2, s3, k1, k2+1, k3+1) {
			return true
		}
	}
	return false
}

func main() {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
	fmt.Println(isInterleave1(s1, s2, s3))
}
