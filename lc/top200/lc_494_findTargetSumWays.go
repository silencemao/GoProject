package main

import (
	"fmt"
)

/*
给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

添加正负号是的数组数字相加减，结果为target

left部分-right部分=target -> left部分 - (sum-left部分)=target -> left部分=(target+sum)/2
即从数组中找出是否存在组合的 和 为left部分。可转化为01背包问题， 类似518题
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 {
		return 0
	}

	target = (sum + target) >> 1 // left部分
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- { // 从后向前遍历，防止nums重复使用
			dp[j] += dp[j-nums[i]] //不能写成 dp[j] = max(dp[j], dp[j-nums[i]]+1)
		}
	}

	if dp[target] == 0 {
		return 0
	}
	return dp[target]
}

/*
深度优先搜索
*/

func help494(nums []int, ind, cur, target int, res *int) {
	if cur == target {
		*res += 1
		//return 此处不能加return
	}
	for i := ind; i < len(nums); i++ {
		cur += nums[i]
		help494(nums, i+1, cur, target, res)
		cur -= nums[i]
	}
}
func findTargetSumWays1(nums []int, target int) int {
	sum := 0
	res := 0
	for _, num := range nums {
		sum += num
	}
	if (target > sum) || (sum+target)%2 == 1 {
		return 0
	}

	target = (sum + target) >> 1 // left部分
	help494(nums, 0, 0, target, &res)
	return res
}

func help494_2(nums []int, target, sum, ind int, res *int) {
	if ind == len(nums) {
		if sum == target {
			*res += 1
		}
		return
	}
	help494_2(nums, target, sum+nums[ind], ind+1, res)
	help494_2(nums, target, sum-nums[ind], ind+1, res)
}

func findTargetSumWays2(nums []int, target int) int {
	sum, res := 0, 0
	help494_2(nums, target, sum, 0, &res)
	return res
}

func main() {
	nums := []int{1, 0}
	target := 1
	fmt.Println(findTargetSumWays(nums, target))
	fmt.Println(findTargetSumWays1(nums, target))
	fmt.Println(findTargetSumWays2(nums, target))
}
