package main

import "fmt"

/*
最短回文串
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
输入：s = "aacecaaa"
输出："aaacecaaa"

在字符串的前部添加，可以将s拆分为t1, t2 最后返回的是 tmp + t1 + t2，三者均可为空串
从上面可以看出t1必须是回文串，tmp是t2的reverse，所以我们的目的是找最长前缀回文
1、l=0, r字符串尾部开始
2、判断 s[l:r]是否为回文
3、将r后面的部分复制出来，翻转
4、加到s头部
*/
func help214(s string, r int) bool {
	for i, j := 0, r; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func shortestPalindrome(s string) string {
	i := 0
	for i = len(s) - 1; i >= 0; i-- {
		if help214(s, i) {
			break
		}
	}
	tmp := []byte(s[i+1:])
	for m, n := 0, len(tmp)-1; m < n; m, n = m+1, n-1 {
		tmp[m], tmp[n] = tmp[n], tmp[m]
	}
	return string(tmp) + s
}

func main() {
	s := "abcd"
	fmt.Println(shortestPalindrome(s))
}
