package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)

/*
有一个无限长的字符串s，...zabcdefghigklmnopqrstuvwxyza...，是连续26个字母的重复组成的
给定一个字符串p，zab 判断字符串中有多少唯一的字符串出现在s中

首先想到的是求出字符串p有多少种组合，以zab为例，z a b za ab zab共六种，因此有6个字符串出现在s中，但是会超时

动态规划的思路，求出以每个字母做为结尾的最大连续子串的长度，然后进行加和即可
以某个字母结尾的连续子串的个数，等于以某个字母结尾的最大连续子串的长度
以zab为例，以b结尾的连续子串有b ab zab 3个
          以b结尾的最大连续子串的长度有为3  二者是相等的
所以求出来以某个字母结尾的最大连续子串的长度，就是以某个字母结尾的连续子串的个数

因此把以26个字母结尾的最大连续子串的长度（有的字母的长度为0）进行求和相加，就是所有子串的个数之和

*/
func findSubstringInWraproundString(p string) int {
	nums := make([]int, 26)
	maxLength := 0
	for i := 0; i < len(p); i++ {
		if i > 0 && ((p[i] - p[i-1] == 1) || (p[i-1] - p[i] == 25)) {
			maxLength++
		} else {
			maxLength = 1
		}
		nums[p[i] - 'a'] = util.MaxInt(nums[p[i] - 'a'], maxLength)
	}

	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
// 给定一个字符串，得到以某个字母结尾的所有子串
func combineStrings(p string) {
	if len(p) == 0 {
		return
	}

	var res []string
	for i := 0; i < len(p); i++ {
		for j := 0; j <= i+1; j++ {
			res = append(res, p[j:i+1])
		}
	}
	fmt.Println(res)
}

func main() {
	p := "zab"
	fmt.Println(findSubstringInWraproundString(p))

	combineStrings(p)
}
