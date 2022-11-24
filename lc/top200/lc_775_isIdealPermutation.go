package main

/*
775. 全局倒置与局部倒置
给你一个长度为 n 的整数数组 nums ，表示由范围 [0, n - 1] 内所有整数组成的一个排列。

全局倒置 的数目等于满足下述条件不同下标对 (i, j) 的数目：
0 <= i < j < n
nums[i] > nums[j]

局部倒置 的数目等于满足下述条件的下标 i 的数目：
0 <= i < n - 1
nums[i] > nums[i + 1]

当数组 nums 中 全局倒置 的数量等于 局部倒置 的数量时，返回 true ；否则，返回 false 。

输入：nums = [1,0,2]
输出：true
解释：有 1 个全局倒置，和 1 个局部倒置。
*/
func isIdealPermutation1(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] && i+1 < j {
				return false
			}
		}
	}
	return true
}

/*
局部倒置 是 i和i+1的关系
全局倒置 是任何i与j的关系
所以局部倒置一定是全局倒置，全局倒置不一定是局部倒置
所以找是否存在一个全局倒置，但不是局部倒置
即 nums[i] > nums[j] i < j-1
*/
func isIdealPermutation(nums []int) bool {
	n := len(nums)
	minSuf := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > minSuf {
			return false
		}
		if minSuf > nums[i] {
			minSuf = nums[i]
		}
	}
	return true
}

func main() {

}
