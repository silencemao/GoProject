package main

import (
	"fmt"
	"strconv"
)

/*
给你一个正整数 num 。你可以交换 num 中 奇偶性 相同的任意两位数字（即，都是奇数或者偶数）。

返回交换 任意 次之后 num 的 最大 可能值。
输入：num = 1234
输出：3412
解释：交换数字 3 和数字 1 ，结果得到 3214 。
交换数字 2 和数字 4 ，结果得到 3412 。
注意，可能存在其他交换序列，但是可以证明 3412 是最大可能值。
注意，不能交换数字 4 和数字 1 ，因为它们奇偶性不同。
*/

func largestInteger(num int) int {
	str := []byte(strconv.Itoa(num))
	//res := ""

	for i := 0; i < len(str); i++ {
		tmp := i
		for j := i + 1; j < len(str); j++ {
			if (str[i]-'0')%2 == (str[j]-'0')%2 && str[j] > str[tmp] {
				tmp = j
			}
		}
		str[i], str[tmp] = str[tmp], str[i]
	}
	//help6037(&res, str, 0)
	x, _ := strconv.Atoi(string(str))
	return x
}

func main() {
	num := 1234
	fmt.Println(largestInteger(num))
}
