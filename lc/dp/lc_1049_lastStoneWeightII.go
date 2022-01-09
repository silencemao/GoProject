package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
有一堆石头，用整数数组stones 表示。其中stones[i] 表示第 i 块石头的重量。

每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为x 和y，且x <= y。那么粉碎的可能结果如下：

如果x == y，那么两块石头都会被完全粉碎；
如果x != y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

拆分成两堆总和最接近的石头，类似第416题
*/
func lastStoneWeightII(stones []int) int {

	res := 0
	for _, v := range stones {
		res += v
	}
	size := res / 2
	dp := make([]int, size+1)
	for _, v := range stones {
		for j := size; j >= v; j-- {
			dp[j] = util.MaxInt(dp[j], dp[j-v]+v)
		}
	}

	return res - dp[size]*2
}
func main() {
	stones := []int{31, 26, 33, 21, 40}
	fmt.Println(lastStoneWeightII(stones))
}
