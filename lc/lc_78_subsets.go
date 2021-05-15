package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var res [][]int
	//var tmp []int

	//dfs90(nums, tmp, &res, 0)
	help78(&res, nums, 0, []int{})
	//help901(nums, 0, []int{}, &res)
	return res
}

func dfs90(nums []int, curr []int, res *[][]int, start int) [][]int {
	currDup := make([]int, len(curr))
	copy(currDup, curr)
	*res = append(*res, currDup)

	for i := start; i < len(nums); i++ {
		curr = append(curr, nums[i])
		*res = dfs90(nums, curr, res, i+1)
		curr = curr[:len(curr)-1]
	}

	return *res
}

func help901(nums []int, idx int, set []int, ans *[][]int) {
	*ans = append(*ans, append([]int{}, set...))

	for i := idx; i < len(nums); i++ {
		set = append(set, nums[i])
		help901(nums, i+1, set, ans)
		set = set[:len(set)-1]
	}
}

func help78(res *[][]int, nums []int, ind int, tmp []int) {
	tmpCur := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		tmpCur[i] = tmp[i]
	}
	*res = append(*res, tmpCur)

	for i := ind; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		help78(res, nums, i+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

// https://leetcode.com/problems/subsets/discuss/27278/C%2B%2B-RecursiveIterativeBit-Manipulation
func subsets1(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {
		size := len(res)
		for j := 0; j < size; j++ {
			res = append(res, append([]int{nums[i]}, res[j]...))
			//res = append(res, res[j])
			//res[len(res)-1] = append(res[len(res)-1], nums[i])
		}
	}

	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsets(nums))
	fmt.Println(subsets1(nums))
}
