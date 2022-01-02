package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。


划分字母区间，遍历字符串，记录每个字母最后一次出现的位置

第二次遍历字符串，找出每个字符最后一次出现的位置，标记为right(right可能>=i)，表示在i左侧所有字符中最后一次出现的位置是right，
当i=right时，i左侧所有的字符都出现在了i左侧，i右侧没有与其相等的字符出现，所以可以分割
长度为right-left
left = right作为新的起点
ababcbacadefegdehijhklij
        此处a最后一次出现，可以进行分割
*/
func partitionLabels(s string) []int {
	set := make(map[byte]int, 0)
	var res []int
	for i := 0; i < len(s); i++ {
		set[s[i]-'a'] = i
	}
	left, right := -1, 0
	for i := 0; i < len(s); i++ {
		right = util.MaxInt(right, set[s[i]-'a'])
		if i == right {
			res = append(res, right-left)
			left = right
		}
	}
	return res
}

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}
