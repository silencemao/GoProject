package main

import "fmt"

/*
最长上升递增子序列，二分查找方法
lastNumOfLis中存储 长度为i的 递增子序列的最大值

遍历数组nums，当元素num>lastNumOfLis[-1]时，则说明num可以作为新的子序列的末尾值（最长上升递增子序列+1）
当num<=lastNumOfLis[-1]时，则覆盖lastNumOfLis中比num大的最小的元素 (更新现有最长上升 递增子序列的最大值)

*/
func lengthOfLIS(nums []int) int {
	lastNumOfLis := make([]int, len(nums)+1)
	lastNumOfLis[1] = nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > lastNumOfLis[l] {
			l += 1
			lastNumOfLis[l] = nums[i]
		} else {
			l, r := 1, l // l要从1开始
			for l < r {
				mid := l + (r-l)>>1
				if lastNumOfLis[mid] < nums[i] { // 找到>nums[i]的最小值
					l = mid + 1
				} else {
					r = mid
				}
			}
			lastNumOfLis[l] = nums[i]
		}
	}
	return l
}

func main() {
	nums := []int{1, 2, -10, -8, -7}
	fmt.Println(lengthOfLIS(nums))
}
