package main

import (
	"fmt"
)

func help491(res *[][]int, tmp, nums []int, ind int) {
	if len(tmp) > 1 {
		*res = append(*res, append([]int{}, tmp...))
	}
	used := make([]int, 201) // nums的范围为[-100, 100]
	for i := ind; i < len(nums); i++ {
		if (len(tmp) > 0 && nums[i] < tmp[len(tmp)-1]) || used[nums[i]+100] == 1 { //同一父节点下的同层上使用过的元素就不能在使用了
			continue
		}

		tmp = append(tmp, nums[i])
		used[nums[i]+100] = 1 //设置完不能重设，在同一根结点下再遇到相同的数字时，就不再使用了
		help491(res, tmp, nums, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

// 给定一个数组（乱序），返回数组中递增子序列的组合，不能重复
//输入：nums = [4,6,7,7]
//输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

func findSubsequences(nums []int) [][]int {
	var res [][]int
	help491(&res, []int{}, nums, 0)
	return res
}
func main() {
	nums := []int{4, 7, 6, 7}
	fmt.Println(findSubsequences(nums))
}
