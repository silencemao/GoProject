package main

import "fmt"

// 中心扩散法 O(n^2)，遍历数组，分别以i、[i:i+1]为中心向两边扩散，获得最长的回文子数组
func longestPalindrome1(s string) string {
	res := ""
	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}

	for i := 0; i < len(s); i++ {
		tmp := validPalindrome(s, i, i)
		if len(tmp) > len(res) {
			res = tmp
		}

		tmp = validPalindrome(s, i, i+1)
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func validPalindrome(s string, left, right int) string {
	bs := []byte(s)
	for left >= 0 && right < len(s) && bs[left] == bs[right] {
		left--
		right++
	}
	return s[left+1 : right] // 此处取值要注意，退出上面的循环，要么left<0 || left位置的值与right位置的不相等 所以left要+1 right同理
}

// 动态规划
// 二维数组dp[i][j]表示从i开始到j结束的子字符串是否为回文
func longestPalindrome2(s string) string {
	res := ""
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	bs := []byte(s)
	for l := 0; l < len(s); l++ { // 子串的长度
		for i := 0; i+l < len(s); i++ { // 起点
			j := i + l // 终点
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 && bs[i] == bs[j] {
				dp[i][j] = 1
			} else if bs[i] == bs[j] && dp[i+1][j-1] == 1 {
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] == 1 && len(bs[i:j+1]) > len(res) {
				res = string(bs[i : j+1])
			}
		}
	}
	return res
}

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome1(s))
	fmt.Println(longestPalindrome2(s))
}
