package main

import (
	"fmt"
	"sort"
)
/*
给你一个 无重复元素 的整数数组candidates 和一个目标整数target，找出candidates中可以使数字和为目标数target 的
所有不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为target 的不同组合数少于 150 个。

1、无重复元素 （不需要记录元素是否使用过）
2、元素可以重复使用 (下一次递归从i开始)
3、组合总和 (判断和是否和target相等)
4、剪枝 (排序+判断当前值与target的关系)

*/
func help39(res *[][]int, candidates, tmp []int, sum, target, ind int) {

	if sum == target {
		*res = append(*res, append([]int{}, tmp...))
	}
	for i := ind; i < len(candidates) && sum + candidates[i] <= target; i++ {

		sum += candidates[i]
		tmp = append(tmp, candidates[i])

		help39(res, candidates, tmp, sum, target, i)

		tmp = tmp[:len(tmp)-1]
		sum -= candidates[i]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	help39(&res, candidates, []int{}, 0, target, 0)
	fmt.Println(res)
	return res
}

func main() {
	candidates := []int{2,7,6,3,5,1}
	target := 9
	combinationSum(candidates, target)
}
