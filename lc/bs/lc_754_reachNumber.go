package main

import "fmt"

/*
在一根无限长的数轴上，你站在0的位置。终点在target的位置。

你可以做一些数量的移动 numMoves :

每次你可以选择向左或向右移动。
第 i次移动（从 i == 1开始，到i == numMoves ），在选择的方向上走 i步。
给定整数target ，返回 到达目标所需的 最小移动次数(即最小 numMoves )。

输入: target = 2
输出: 3
解释:
第一次移动，从 0 到 1 。
第二次移动，从 1 到 -1 。
第三次移动，从 -1 到 2 。


p+n=s
p-n=target
s-target=2*n>=0
s=(1+i)*i/2
(1+i)*i/2-target为偶数且大于等于0
*/
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	s, i := 0, 1
	for ; i < 1<<31-1; i++ {
		s = (1 + i) * i / 2
		if (s-target)%2 == 0 && s >= target {
			return i
		}
	}
	return -1
}
func main() {
	target := 1000
	fmt.Println(reachNumber(target))
}
