package main

import (
	"fmt"
	"math"
)

/*
返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。
请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0是有效的。

输入：n = 3, k = 7
输出：[181,292,707,818,929]
解释：注意，070 不是一个有效的数字，因为它有前导零。
*/
func help967(res *[]int, tmp []int, n, k int) {
	if len(tmp) == n {
		//fmt.Println(tmp)
		t := 0
		for _, m := range tmp {
			t = t*10 + m
		}
		*res = append(*res, t)
		return // 需要返回 否则会溢出
	}
	for i := 0; i < 10; i++ {
		if len(tmp) > 0 && (tmp[0] == 0 || int(math.Abs(float64(tmp[len(tmp)-1]-i))) != k) {
			continue
		}
		tmp = append(tmp, i)
		help967(res, tmp, n, k)
		tmp = tmp[:len(tmp)-1]
	}
}

func numsSameConsecDiff(n int, k int) []int {
	var res []int
	help967(&res, []int{}, n, k)
	return res
}

func main() {
	n, k := 3, 7
	fmt.Println(numsSameConsecDiff(n, k))
}
