package main

import "strings"

/*检查两个字符串数组是否相等
给你两个字符串数组 word1 和 word2 。如果两个数组表示的字符串相同，返回 true ；否则，返回 false 。

数组表示的字符串是由数组中的所有元素 按顺序 连接形成的字符串。

输入：word1 = ["ab", "c"], word2 = ["a", "bc"]
输出：true
解释：
word1 表示的字符串为 "ab" + "c" -> "abc"
word2 表示的字符串为 "a" + "bc" -> "abc"
两个字符串相同，返回 true

*/
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := strings.Join(word1, "")
	s2 := strings.Join(word2, "")
	return s1 == s2
}
func main() {

}
