package main

import (
	"fmt"
	"strconv"
)

/*
给定一个非负整数N，找出小于或等于N的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
（当且仅当每个相邻位数上的数字x和y满足x <= y时，我们称这个整数是单调递增的。）
例如：n=332
返回299

1、数字转换称字符串
2、从后向前遍历，遇到ss[i] > ss[i+1]
3、ss[i]-1，[i+1, len(s)-1]全部变为9
4、注意byte 和 数字的区别
*/
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)

	ss := []byte(s)
	for i := len(ss) - 2; i >= 0; i-- {
		if ss[i] > ss[i+1] {
			ss[i] = ss[i] - 1
			for j := i + 1; j < len(ss); j++ {
				ss[j] = '9'
			}
		}
	}
	n, _ = strconv.Atoi(string(ss))
	return n
}

func main() {
	s := 332
	fmt.Println(monotoneIncreasingDigits(s))
}
