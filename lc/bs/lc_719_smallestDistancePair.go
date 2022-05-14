package main

import (
	"fmt"
	"sort"
)

/*
找出第k小的距离对
给定一个整数数组，返回所有数对之间的第 k 个最小距离。一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。
输入：
nums = [1,3,1]
k = 1
输出：0
解释：
所有数对如下：
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
因此第 1 个最小距离的数对是 (1,1)，它们之间的距离为 0。

*/
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, nums[len(nums)-1]-nums[0]
	for l < r {
		mid := l + (r-l)>>1
		cnt := help719(nums, mid)
		if cnt < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func help719(nums []int, target int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ { //固定i
		j := len(nums) - 1
		for ; j > i; j-- { // 找到第一个nums[j]-nums[i]<=target的元素，i和[i+1,j]之间任一元素组合都<=target，共计j-i个组合
			if nums[j]-nums[i] <= target {
				break
			}
		}
		cnt += j - i
	}
	return cnt
}

func main() {

	nums := []int{38, 33, 57, 65, 13, 2, 86, 75, 4, 56}
	k := 26
	fmt.Println(smallestDistancePair(nums, k))
}
