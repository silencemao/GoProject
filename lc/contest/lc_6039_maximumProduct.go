package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
. K 次增加后的最大乘积
给你一个非负整数数组nums和一个整数k。每次操作，你可以选择nums中 任一元素并将它 增加1。
请你返回 至多k次操作后，能得到的nums的最大乘积。由于答案可能很大，请你将答案对109 + 7取余后返回。

输入：nums = [0,4], k = 5
输出：20
解释：将第一个数增加 5 次。
得到 nums = [5, 4] ，乘积为 5 * 4 = 20 。
可以证明 20 是能得到的最大乘积，所以我们返回 20 。
*/

func maximumProduct(nums []int, k int) int {
	h := hp{nums}
	for heap.Init(&h); k > 0; k-- {
		h.IntSlice[0]++ // 每次给最小的加一
		heap.Fix(&h, 0)
	}
	fmt.Println(nums)
	ans := 1
	for _, num := range nums {
		ans = ans * num % (1e9 + 7)
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

func maximumProduct1(nums []int, k int) int {
	res := 1

	for k > 0 {
		res = 1
		sort.Ints(nums)
		for i := range nums {
			if i == 0 {
				nums[i] += 1
			}
			res = res * nums[i] % (1e9 + 7)
		}
		k--
	}
	return res
}

func main() {
	nums := []int{0, 4}
	k := 5
	fmt.Println(maximumProduct(nums, k))
	fmt.Println(maximumProduct1(nums, k))
}
