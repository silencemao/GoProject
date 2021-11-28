package main

import (
	"fmt"
	"sort"
)

/*
前K个高频元素
此题还可以用堆来做，学习中
*/
func topKFrequent(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}
	set := make(map[int]int, 0)
	for _, num := range nums {
		set[num] += 1
	}
	tmp := []int{}
	for num := range set {
		tmp = append(tmp, num)
	}
	sort.SliceStable(tmp, func(i, j int) bool {
		return set[tmp[i]] > set[tmp[j]]
	})

	return tmp[:k]
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
