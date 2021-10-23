package main

import (
	"fmt"
	"sort"
)

func help40(res *[][]int, candidates, used, tmp []int, ind, sum, target int) {
	if sum == target {
		*res = append(*res, append([]int{}, tmp...))
		return
	}

	for i := ind; i < len(candidates) && candidates[i]+sum <= target; i++ {
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == 0 {
			continue
		}
		sum += candidates[i]
		tmp = append(tmp, candidates[i])
		used[i] = 1
		help40(res, candidates, used, tmp, i+1, sum, target)
		sum -= candidates[i]
		tmp = tmp[:len(tmp)-1]
		used[i] = 0
	}
}

/*
给定数组candidates，以及目标数字target，数组中存在重复，返回和为target的数字组合，组合之间不能重复，candidates 中的每个数字在每个组合中只能使用一次。
输入: candidates =[10,1,2,7,6,1,5], target =8,
输出:
[ [1,1,6], [1,2,5],[1,7],[2,6] ]
组合不能重复，组合内可以使用candidates中的重复数字 比如上面例子中的1
那如何避免组合之间重复呢？
1、先排序
2、也就是当i>0 & candidates[i]==candidates[i-1]时，需要进行过滤，但是仅凭前两个条件会过滤掉[1,1,6]这种的组合，
3、所以还需要一个布尔数组used来记录数字的使用情况，
	若在构成一个组合的过程中，第一个1使用了，标记为true，遇到第二个1时，used[i-1]=true，第二个1也可以使用
	若在构成不同组合时，当遇到第二个1时，观察used[i-1]的情况，若前一个1在本次组合内没有使用，若本次组合使用第二个1，则会生成和使用第一个1同样的组合，存在重复
4、所以过滤条件再加一条 used[i-1]=false，表示若生成当前组合则会存在重复，所以需要continue
*/
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	used := make([]int, len(candidates))
	sort.Ints(candidates)
	help40(&res, candidates, used, []int{}, 0, 0, target)
	for _, r := range res {
		fmt.Println(r)
	}
	return res
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
