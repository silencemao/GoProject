package main

/*
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。
输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。

*/
func findUnsortedSubarray(nums []int) int {
	l, r, sz := -1, -1, len(nums)
	minNum, maxNum := 1<<31-1, -1<<31-1

	for i := 0; i < len(nums); i++ {
		if nums[i] >= maxNum {
			maxNum = nums[i]
		} else {
			r = i
		}

		if nums[sz-1-i] <= minNum {
			minNum = nums[sz-1-i]
		} else {
			l = sz - 1 - i
		}
	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}
func main() {

}
