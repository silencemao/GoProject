package main

import "fmt"

/*
给定一个字符串s，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1。

输入: s = "leetcode"
输出: 0

*/
func firstUniqChar(s string) int {
	dp := make(map[byte]int, 0)
	for i := 'a'; i < 'z'; i++ {
		dp[byte(i)] = -1
	}
	for i := 0; i < len(s); i++ {
		if val, ok := dp[s[i]]; ok {
			if val == -1 {
				dp[s[i]] = i
			} else {
				dp[s[i]] = 100
			}
		} else {
			dp[s[i]] = i
		}
	}
	res := 1<<31 - 1
	for i := 'a'; i < 'z'; i++ {
		if val, ok := dp[byte(i)]; ok {
			if val != -1 && val != 100 {
				if val < res {
					res = val
				}
			}
		}
	}
	return res
}

func firstUniqChar1(s string) int {
	var dp [26]int
	//n := len(s)
	for _, c := range s {
		dp[c-'a'] += 1
	}
	for i, c := range s {
		if dp[c-'a'] == 1 {
			return i
		}
	}
	return -1

}

func firstUniqChar2(s string) int {
	flag := make([]int, 26)
	n := len(s) + 1
	for i := range flag {
		flag[i] = n
	}
	res := 1<<31 - 1
	for i, c := range s {
		ind := int(c - 'a')
		if flag[ind] == n {
			flag[ind] = i
		} else {
			flag[ind] = n + 1
		}
	}
	for _, ind := range flag {
		if ind < res {
			res = ind
		}
	}
	if res == 1<<31-1 || res >= n {
		res = -1
	}
	return res

}

func main() {
	s := "leetcodeabcdtel"
	fmt.Println(firstUniqChar1(s))
	fmt.Println(firstUniqChar2(s))
}
