package main

import (
	"sort"
)

/*
给定一个无重复元素的正整数数组candidates和一个正整数target，找出candidates中所有可以使数字和为目标数target的唯一组合。
candidates中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
对于给定的输入，保证和为target 的唯一组合数少于 150 个。
*/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	help39(&res, []int{}, candidates, target, 0, 0)
	//for _, r := range res {
	//	fmt.Println(r)
	//}
	return res
}

func help39(res *[][]int, tmp, candidates []int, target int, ind, sum int) {
	if sum == target {
		*res = append(*res, append([]int{}, tmp...))
		return
	}

	for i := ind; i < len(candidates) && candidates[i]+sum <= target; i++ {
		tmp = append(tmp, candidates[i])
		sum += candidates[i]

		help39(res, tmp, candidates, target, i, sum) // 此处使用i表示可以重复使用

		tmp = tmp[:len(tmp)-1]
		sum -= candidates[i]
	}
}

func main() {
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9
	combinationSum(candidates, target)
}
