package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
最后一块石头的重量
有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
每一回合任意选取两个石头，石头重量为x，y，进行打击，x变成0，y变成abs(y-x)，最后数组中最多剩余一块石头，返回石头的最小可能重量。

1、若只有两个石头，返回石头最小可能的重量，两个石头的重量越接近，则返回石头的重量越小
2、若有三个石头，则先让两个最大的石头打击，最后剩余重量再和小的石头打击，剩余石头的重量最小 [1, 4, 5]
4、若有四个石头，则两个最大的互相打击，两个最小的互相打击，余下的重量再互相打击[1, 2, 7, 9]
5、根据上面的描述 我们可以递推出若互相打击的两个石头分别分到A B两组，则A B内重量之和越接近，则最后剩余的重量则越小
6、此题是不是和第416题有点接近了

其实此题就是找出接近 所有石头重量之和/2的 的石头组合，即找出石头的组合，其重量最接近sum/2的石头重量，
是不是可以转换为01背包问题了
那我们的背包大小为sum/2 物品就是stones

最后剩余最小的石头重量就是abs(sum-2*target)
*/
func lastStoneWeightII(stones []int) int {
	res, sum := 0, 0

	for _, num := range stones {
		sum += num
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 0
	for _, num := range stones {
		for j := target; j >= num; j-- {
			dp[j] = util.MaxInt(dp[j], dp[j-num]+num)
		}
	}
	res = sum - 2*dp[target]
	if res < 0 {
		return -res
	}
	return res
}
func main() {
	stones := []int{31, 26, 33, 21, 40}
	fmt.Println(lastStoneWeightII(stones))
}
