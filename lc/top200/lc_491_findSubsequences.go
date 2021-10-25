package main

import (
	"fmt"
)

func help491(res *[][]int, tmp, nums, used []int, ind int) {
	if len(tmp) > 1 {
		*res = append(*res, append([]int{}, tmp...))
	}
	for i := ind; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 { // 避免结果重复
			continue
		}
		if len(tmp) == 0 || (len(tmp) > 0 && nums[i] >= tmp[len(tmp)-1]) { // 判断是否会成为递增子数组
			tmp = append(tmp, nums[i])
			used[i] = 1
			help491(res, tmp, nums, used, i+1)
			used[i] = 0
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func findSubsequences(nums []int) [][]int {
	var res [][]int
	used := make([]int, len(nums))
	help491(&res, []int{}, nums, used, 0)
	return res
}
func main() {
	nums := []int{4, 4, 3, 2, 1}
	fmt.Println(findSubsequences(nums))
}
