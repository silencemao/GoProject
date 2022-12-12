package main

/*
所有子字符串的美丽值之和
一个字符串的 美丽值定义为：出现频率最高字符与出现频率最低字符的出现次数之差。

比方说，"abaacc"的美丽值为3 - 1 = 2。
给你一个字符串s，请你返回它所有子字符串的美丽值之和。

输入：s = "aabcb"  输出：5
解释：美丽值不为零的字符串包括 ["aab","aabc","aabcb","abcb","bcb"] ，每一个字符串的美丽值都为 1 。
*/
func beautySum(s string) int {
	res := 0
	for i := range s {
		cnt := [26]int{}
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a'] += 1
			mx, mi := 0, 1000
			for _, v := range cnt {
				if v <= 0 {
					continue
				}
				if mx < v {
					mx = v
				}
				if mi > v {
					mi = v
				}
			}
			res += mx - mi
		}
	}
	return res
}

func main() {

}
