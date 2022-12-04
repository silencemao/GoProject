package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
给定两个数组和一个目标
baseCosts中的元素必须且只能选一个
toppingCosts中的元素可以选一个或多个，每个元素最多选两次
target 目标值
返回最接近目标值的 组合的 和，如有两个和目标值距离一样的结果，返回较小的那一个
*/
func help1774(toppingCosts []int, target int, sum int, res *int, ind int) {
	if util.Abs(sum-target) < util.Abs(*res-target) {
		*res = sum
	} else if util.Abs(sum-target) == util.Abs(*res-target) {
		if sum < *res {
			*res = sum
		}
	}
	for i := ind; i < len(toppingCosts); i++ {
		if util.Abs(sum+toppingCosts[i]-target) > util.Abs(sum-target) {
			continue
		}
		sum += toppingCosts[i]
		help1774(toppingCosts, target, sum, res, i+1)
		sum -= toppingCosts[i]
	}
}

/*
dfs 遍历baseCosts 深度搜索toppingCosts
*/
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	res := baseCosts[0]
	toppingCosts = append(toppingCosts, toppingCosts...)
	sort.Ints(toppingCosts)
	for i := 0; i < len(baseCosts); i++ {
		help1774(toppingCosts, target, baseCosts[i], &res, 0)
	}
	return res
}

/*
动态规划
*/
func closestCost2(baseCosts []int, toppingCosts []int, target int) int {
	toppingCosts = append(toppingCosts, toppingCosts...)

	dp := make([]bool, 20000+1)
	for _, b := range baseCosts {
		dp[b] = true
	}
	for _, t := range toppingCosts {
		for i := len(dp) - 1; i >= t; i-- {
			dp[i] = dp[i] || dp[i-t]
		}
	}
	res := 1<<31 - 1
	for i := range dp {
		if dp[i] && util.Abs(res-target) > util.Abs(i-target) {
			res = i
		}
	}
	return res
}

func main() {
	baseCosts := []int{10}
	toppingCosts := []int{1}
	target := 1
	fmt.Println(closestCost(baseCosts, toppingCosts, target))
}
