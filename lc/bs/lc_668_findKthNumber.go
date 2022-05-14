package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
乘法表中第k小的数
几乎每一个人都用乘法表。但是你能在乘法表中快速找到第k小的数字吗？
给定高度m、宽度n 的一张m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。

输入: m = 3, n = 3, k = 5
输出: 3
解释:
乘法表:
1	2	3
2	4	6
3	6	9
第5小的数字是 3 (1, 2, 2, 3, 3).

*/
func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	for l < r {
		mid := l + (r-l)>>1
		pos := help668(m, n, mid)

		if pos < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func help668(m, n, target int) int {
	cnt := 0
	for i := 1; i <= m && i <= target; i++ {
		cnt += util.MinInt(target/i, n)
	}
	return cnt
}

func main() {

	m, n, k := 7341, 13535, 12330027
	fmt.Println(findKthNumber(m, n, k))
}
