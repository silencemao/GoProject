package main

import (
	"fmt"
)
// 含有重复数字的全排列，找出不重复的全排列
func dfs2(nums *[]int, res *[][]int, index int) {
	if index >= len(*nums) {
		*res = append(*res, append([]int{}, *nums...))
		return
	}

	set := make(map[int]int)
	for i := index; i < len(*nums); i++ {
		//if i > 0 && (*nums)[i] == (*nums)[i-1] { 这种写法是错误的，在当前case中因为直接错过了，index=0 i=1的情况
		//	continue
		//}
		if _, ok := set[(*nums)[i]]; !ok {
			set[(*nums)[i]] = 1
			(*nums)[i], (*nums)[index] = (*nums)[index], (*nums)[i]
			dfs2(nums, res, index+1)
			(*nums)[index], (*nums)[i] = (*nums)[i], (*nums)[index]
		} else {
			fmt.Println(ok, res, nums)
		}
	}
}

func permutationUnique(nums []int) [][]int {
	var res [][]int
	dfs2(&nums, &res, 0)
	return res
}

//func main() {
//	var nums = []int{1, 1, 2}
//	res := permutationUnique(nums)
//	fmt.Println(res)
//}