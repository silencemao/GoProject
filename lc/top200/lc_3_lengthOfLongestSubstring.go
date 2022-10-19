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

func lengthOfLongestSubstring2(s string) int {
	runeMap := make(map[rune]int)
	max, beginAt := 0, 0
	for k, v := range s {
		if value, ok := runeMap[v]; ok && value >= beginAt {
			beginAt = value + 1
		}
		if k-beginAt+1 > max {
			max = k - beginAt + 1
		}
		runeMap[v] = k
	}
	return max
}

func lengthOfLongestSubstring3(s string) int {
	l, r, res := -1, 0, 0
	set := map[byte]int{}
	for ; r < len(s); r++ {
		if ind, ok := set[s[r]]; ok {
			l = util.MaxInt(ind, l)
		}
		res = util.MaxInt(r-l, res)
		set[s[r]] = r
	}
	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring1(s))
}
