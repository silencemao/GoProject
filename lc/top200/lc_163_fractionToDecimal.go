package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"strconv"
	"strings"
)

/*
给定两个整数，分别表示分数的分子numerator 和分母 denominator，以 字符串形式返回小数 。

如果小数部分为循环小数，则将循环的部分括在括号内。

如果存在多个答案，只需返回 任意一个 。

对于所有给定的输入，保证 答案字符串的长度小于 104 。。
*/
func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	var res []string
	if numerator*denominator < 0 {
		res = append(res, "-")
	}
	set := make(map[int]int, 0)

	a, b := util.Abs(numerator), util.Abs(denominator)
	res = append(res, strconv.Itoa(a/b))
	res = append(res, ".")
	a %= b
	for a != 0 {
		set[a] = len(res)
		a *= 10
		res = append(res, strconv.Itoa(a/b))
		a %= b
		if ind, ok := set[a]; ok {
			fmt.Println(res, ind)
			//r := res[0] + "." + strings.Join(res[1:ind], "") + "(" + strings.Join(res[ind:], "") + ")"
			r := strings.Join(res[:ind], "") + "(" + strings.Join(res[ind:], "") + ")"
			return r
		}
	}
	return strings.Join(res, "")
}

func main() {
	a, b := -1, 6
	fmt.Println(fractionToDecimal(a, b))
}
