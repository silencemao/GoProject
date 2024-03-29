package main

import "fmt"

/*
你总共有 n 枚硬币，并计划将它们按阶梯状排列。对于一个由 k 行组成的阶梯，其第 i 行必须正好有 i 枚硬币。阶梯的最后一行 可能 是不完整的。

给你一个数字 n ，计算并返回可形成 完整阶梯行 的总行数。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/arranging-coins
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func arrangeCoins(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)>>1
		tmp := (1 + mid) * mid / 2
		if tmp <= n {
			if tmp+mid+1 > n {
				return mid
			}
			if tmp+mid+1 == n {
				return mid + 1
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func main() {
	n := 10
	fmt.Println(arrangeCoins(n))
}
