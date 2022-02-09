package main

import "fmt"

/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
*/

func help90(res *[][]int, nums, tmp []int, used map[int]bool, ind int) {
	*res = append(*res, append([]int{}, tmp...))

	for i := ind; i < len(nums); i++ {
		// 举例: i=2, 在本次循环中nums[1]没有使用，说明是nums[1]和nums[2]不会同时出现，而是作为独立的个体出现在不同的子集中，
		// 所以为避免重复只使用nums[1],nums[2]不再使用
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		used[i] = true
		tmp = append(tmp, nums[i])
		help90(res, nums, tmp, used, i+1)
		used[i] = false
		tmp = tmp[:len(tmp)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	used := make(map[int]bool, 0)
	help90(&res, nums, []int{}, used, 0)
	return res
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
