package main

import "sort"

/*
火柴拼正方形
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i个火柴棒的长度。
你要用 所有的火柴棍拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。

如果你能使这个正方形，则返回 true ，否则返回 false 。
同第698

相当于数组的元素是否可以拆成4个子数组，每个子数组的和相等
*/
func makesquare(matchsticks []int) bool {
	sum, k := 0, 4
	if len(matchsticks) < 4 {
		return false
	}
	sort.Ints(matchsticks)
	for _, num := range matchsticks {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	used := make([]bool, len(matchsticks))
	return help(0, 0, sum/k, k, matchsticks, used)
}

func help(ind, curSum, target, k int, nums []int, used []bool) bool {
	if k == 1 {
		return true
	}
	if curSum == target {
		return help(0, 0, target, k-1, nums, used)
	}
	for i := ind; i < len(nums) && curSum+nums[i] <= target; i++ {
		if used[i] {
			continue
		}
		curSum += nums[i]
		used[i] = true

		if help(i+1, curSum, target, k, nums, used) {
			return true
		}

		curSum -= nums[i]
		used[i] = false
	}
	return false
}

func main() {

}
