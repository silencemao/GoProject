package main

import (
	"fmt"
	"sort"
)

/*
和第40、78题类似
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

*/
func help90(res *[][]int, nums, tmp, used []int, ind int) {
	*res = append(*res, append([]int{}, tmp...))
	if ind >= len(nums) {
		return
	}
	for i := ind; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
			continue
		}
		tmp = append(tmp, nums[i])
		used[i] = 1
		help90(res, nums, tmp, used, i+1)
		used[i] = 0
		tmp = tmp[:len(tmp)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	used := make([]int, len(nums))
	help90(&res, nums, []int{}, used, 0)

	return res
}

//同类型 78题
func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
