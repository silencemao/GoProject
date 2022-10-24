package main

import "GoProject/leetcode/util"

/*
给定一个数组nums，将其划分为两个连续子数组left和right，使得：
left中的每个元素都小于或等于right中的每个元素。
left 和right都是非空的。
left 的长度要尽可能小。
在完成这样的分组后返回left的长度。

用例可以保证存在这样的划分方法。

示例 1：

输入：nums = [5,0,3,8,6]
输出：3
解释：left = [5,0,3]，right = [8,6]

*/
func partitionDisjoint(nums []int) int {
	n := len(nums)
	minRight := make([]int, n)
	minRight[n-1] = nums[n-1] // 存储i-n区间内的最小值
	for i := n - 2; i >= 0; i-- {
		minRight[i] = util.MinInt(minRight[i+1], nums[i])
	}
	maxLeft := nums[0] // 0-i之间的最大值
	for i := 1; i < n; i++ {
		if maxLeft <= minRight[i] { // 左边最大值<=右边最小值
			return i
		}
		maxLeft = util.MaxInt(nums[i], maxLeft)
	}
	return n
}

func main() {

}
