package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	dp := make([]int, 256)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	maxLength := 0
	lastPos := -1
	for i := 0; i < len(s); i++ {
		if dp[s[i]] != -1 && lastPos < dp[s[i]] {
			lastPos = dp[s[i]]
		}
		if i-lastPos > maxLength {
			maxLength = i - lastPos
		}
		dp[s[i]] = i
	}

	return maxLength
}

/*
滑动窗口
*/
func lengthOfLongestSubstring1(s string) int {
	res, l := 0, 0
	window := make(map[byte]int, 128)
	for r := 0; r < len(s); r++ {
		window[s[r]]++

		for window[s[r]] > 1 {
			window[s[l]]--
			l++
		}
		res = util.MaxInt(res, r-l+1)
	}

	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring1(s))
}
