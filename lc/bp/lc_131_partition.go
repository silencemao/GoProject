package main

import "fmt"

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

一个字符串拆成多个回文子串，
*/
func isPartition(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func help131(s string, res *[][]string, tmp []string, ind int) {
	if ind >= len(s) {
		*res = append(*res, append([]string{}, tmp...))
	}
	for i := ind; i < len(s); i++ {
		if !isPartition(s, ind, i) {
			continue
		}
		tmp = append(tmp, s[ind:i+1])
		help131(s, res, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func partition(s string) [][]string {
	var res [][]string
	help131(s, &res, []string{}, 0)
	return res
}

func help131II(s string, res *[][]string, tmp []string, ind int, dp [][]int) {
	if ind >= len(s) {
		*res = append(*res, append([]string{}, tmp...))
	}
	for i := ind; i < len(s); i++ {
		if dp[ind][i]==0 {
			continue
		}
		tmp = append(tmp, s[ind:i+1])
		help131II(s, res, tmp, i+1, dp)
		tmp = tmp[:len(tmp)-1]
	}
}

func partitionII(s string) [][]string {
	var res [][]string
	size := len(s)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}
	for i := size-1; i >= 0; i-- {
		for j := i+1; j < size; j++ {
			if s[i] == s[j] && ((j-i <= 2) || dp[i+1][j-1]==1) {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	help131II(s, &res, []string{}, 0, dp)
	return res
}


func main() {
	s := "abaamnm"
	fmt.Println(partition(s))
	fmt.Println(partitionII(s))
}
