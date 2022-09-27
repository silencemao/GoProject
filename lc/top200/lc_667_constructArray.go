package main

import "fmt"

/*
给你两个整数 n 和 k ，请你构造一个答案列表 answer ，该列表应当包含从 1 到 n 的 n 个不同正整数，并同时满足下述条件：

假设该列表是 answer =[a1, a2, a3, ... , an] ，那么列表 [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数。
返回列表 answer 。如果存在多种答案，只需返回其中 任意一种 。

输入：n = 3, k = 1
输出：[1, 2, 3]
解释：[1, 2, 3] 包含 3 个范围在 1-3 的不同整数，并且 [1, 1] 中有且仅有 1 个不同整数：1

*/

func constructArray(n int, k int) []int {
	var res []int
	l, r := 1, k+1
	for l < r {
		res = append(res, l)
		res = append(res, r)
		l, r = l+1, r-1
	}
	if l == r {
		res = append(res, l)
	}
	for i := len(res); i < n; i++ {
		res = append(res, i+1)
	}
	fmt.Println(res)
	return res
}

func main() {

}
