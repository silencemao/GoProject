package main

import "fmt"
/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/
func help78(nums []int, res *[][]int, tmp []int, ind int) {
	//if ind >= len(nums) {
	*res = append(*res, append([]int{}, tmp...))
		//return
	//}

	for i := ind; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		help78(nums, res, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func subsets(nums []int) [][]int {
	var res [][]int
	help78(nums, &res, []int{}, 0)
	return res
}

func subsets2(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for _, num := range nums {
		for _, arr := range res {
			tmp := append([]int{num}, arr...)
			res = append(res, tmp)
		}
		//res = append(res, []int{num})
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
	fmt.Println(subsets2(nums))
}
