package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
最大平均值和的分组

给定数组nums和一个整数k。我们将给定的数组nums分成 最多k个相邻的非空子数组 。分数 由每个子数组内的平均值的总和构成。
注意我们必须使用 nums 数组中的每一个数进行分组，并且分数不一定需要是整数。
返回我们所能得到的最大 分数 是多少。答案误差在10-6内被视为是正确的。

输入: nums = [9,1,2,3,9], k = 3
输出: 20.00000
解释:
nums 的最优分组是[9], [1, 2, 3], [9]. 得到的分数是 9 + (1 + 2 + 3) / 3 + 9 = 20.
我们也可以把 nums 分成[9, 1], [2], [3, 9].
这样的分组得到的分数为 5 + 2 + 6 = 13, 但不是最大值.

1、前缀和
2、dp
3、分情况讨论 (i个数分成j份 份j=1和j>1两种情况)
j>1时，计算第j个子数组的平均值 + 前j-1个子数组的平均值之和
*/
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ { // 前缀和
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	dp := make([][]float64, n+1) // dp[i][j] i个元素拆分为j个子数组的平均值之和
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j < i+1 && j < k+1; j++ { // 分成j份
			if j == 1 {
				dp[i][j] = float64(preSum[i]) / float64(i)
			} else {
				for m := 2; m <= i; m++ { //最后一个子数组的元素索引起始位置
					dp[i][j] = util.MaxFloat64(dp[i][j], dp[m-1][j-1]+float64(preSum[i]-preSum[m-1])/float64(i-m+1))
				}
			}
		}
	}
	return dp[n][k]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 4
	fmt.Println(largestSumOfAverages(nums, k))
}
