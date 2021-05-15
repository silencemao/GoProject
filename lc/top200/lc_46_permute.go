package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	permuteHelp(&res, nums, 0)
	for _, r := range res {
		fmt.Println(r)
	}

	return res
}

func permuteHelp(res *[][]int, nums []int, ind int) {
	if ind >= len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	}
	for i := ind; i < len(nums); i++ { // 从ind开始
		nums[i], nums[ind] = nums[ind], nums[i]
		permuteHelp(res, nums, ind+1)
		nums[i], nums[ind] = nums[ind], nums[i]
	}
}

func permute1(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	permuteHelp1(&res, nums, []int{}, used, 0)
	return res
}

func permuteHelp1(res *[][]int, nums []int, tmp []int, used []bool, ind int) {
	if ind >= len(nums) {
		fmt.Println(tmp)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ { // 从0开始 因为有used控制 每个是否出现过
		if !used[i] {
			tmp = append(tmp, nums[i])
			used[i] = true

			permuteHelp1(res, nums, tmp, used, ind+1)

			used[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	permute1(nums)
}
