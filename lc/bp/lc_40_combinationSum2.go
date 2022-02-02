package main

import (
	"fmt"
	"sort"
)

/*
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。

1、组合总和
2、有重复数字 (标记是否使用过)
3、一个数字不能重复使用，相同的数字在一个组合内可以重复使用 (同上)
4、剪枝
对比第39题

*/

func help40(res *[][]int, candidates, tmp []int, sum, target, ind int, used map[int]bool) {
	if sum == target {
		*res = append(*res, append([]int{}, tmp...))
		return
	}
	for i := ind; i < len(candidates) && sum + candidates[i] <= target; i++ {
		if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		sum += candidates[i]
		tmp = append(tmp, candidates[i])

		help40(res, candidates, tmp, sum, target, i+1, used)

		used[i] = false
		sum -= candidates[i]
		tmp = tmp[:len(tmp)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)

	used := make(map[int]bool, 0)
	help40(&res, candidates, []int{}, 0, target, 0, used)
	fmt.Println(res)

	return res
}


func main() {
	candidates := []int{2,5,2,1,2}
	target := 5
	combinationSum2(candidates, target)
}
