package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
https://leetcode-cn.com/problems/permutations-ii/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liwe-2/
*/
func help47(res *[][]int, nums, tmp, used []int) {
	if len(tmp) == len(nums) {
		*res = append(*res, append([]int{}, tmp...))
		return
	}
	for i := 0; i < len(nums); i++ {
		// 1 1' 表示第一个1和 第二个1
		// 先选择1，再选择1' 满足条件，保留结果
		// 回退到选择1之前
		// 先选择1'，再选择1 结果和第一次的重复，剪枝掉
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 { // 1 和 1'相等，
			continue
		}
		if used[i] == 0 {
			tmp = append(tmp, nums[i])
			used[i] = 1
			help47(res, nums, tmp, used)
			tmp = tmp[:len(tmp)-1]
			used[i] = 0
		}
	}
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	used := make([]int, len(nums))
	help47(&res, nums, []int{}, used)
	return res
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
