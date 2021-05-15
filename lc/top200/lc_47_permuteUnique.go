package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	permuteUniqueHelp(&res, nums, 0)
	for _, r := range res {
		fmt.Println(r)
	}

	return res
}

func permuteUniqueHelp(res *[][]int, nums []int, ind int) {
	if ind >= len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	}
	set := make(map[int]int)
	for i := ind; i < len(nums); i++ { // 从ind开始
		if _, ok := set[nums[i]]; !ok {
			set[nums[i]] = 1
			nums[i], nums[ind] = nums[ind], nums[i]
			permuteUniqueHelp(res, nums, ind+1)
			nums[i], nums[ind] = nums[ind], nums[i]
		}
	}
}

func main() {
	nums := []int{1, 1, 2}
	permuteUnique(nums)
}
