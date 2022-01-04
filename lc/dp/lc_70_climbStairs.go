package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

斐波那契数列变体
*/
func climbStairs(n int) int {
	res := 0

	if n <= 2 {
		return n
	}
	f, s := 1, 2
	for i := 3; i <= n; i++ {
		res = f + s
		f, s = s, res
	}

	return res
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
