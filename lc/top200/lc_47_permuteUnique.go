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

func h31(res *[][]int, nums, used, tmp []int, ind int) {
	if len(tmp) == len(nums) {
		*res = append(*res, append([]int{}, tmp...))
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == 0 {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				continue
			}
			tmp = append(tmp, nums[i])
			used[i] = 1

			h31(res, nums, used, tmp, i+1)

			tmp = tmp[:len(tmp)-1]
			used[i] = 0
		}
	}
}

func f31(nums []int) [][]int {
	var res [][]int
	used := make([]int, len(nums))
	h31(&res, nums, used, []int{}, 0)
	fmt.Println(res)
	return res
}

func main() {
	nums := []int{1, 1, 2}
	permuteUnique(nums)
	f31(nums)
}
