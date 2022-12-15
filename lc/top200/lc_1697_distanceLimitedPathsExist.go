package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
	"sort"
)

/*
给你一个 n个点组成的无向图边集edgeList，其中edgeList[i] = [ui, vi, disi]表示点ui 和点vi之间有一条长度为disi的边。
请注意，两个点之间可能有 超过一条边。

给你一个查询数组queries，其中queries[j] = [pj, qj, limitj]，你的任务是对于每个查询queries[j]，判断是否存在从pj到qj的路径，
且这条路径上的每一条边都 严格小于limitj。

请你返回一个 布尔数组answer，其中answer.length == queries.length，当queries[j]的查询结果为true时，answer 第j个值为true，否则为false。

*/
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	ans := make([]bool, len(queries))

	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	for i := range queries { // 保持原有queries的顺序
		queries[i] = append(queries[i], i)
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	un := _struct.NewUnionFind(len(edgeList))

	k := 0
	for _, q := range queries {
		for ; k < len(edgeList) && edgeList[k][2] < q[2]; k++ {
			un.Union(edgeList[k][0], edgeList[k][1])
		}
		ans[q[3]] = un.Find(q[0]) == un.Find(q[1])
	}

	return ans
}

func main() {
	n := 3
	edgeList := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	queries := [][]int{{0, 1, 2}, {0, 2, 5}}
	fmt.Println(distanceLimitedPathsExist(n, edgeList, queries))

}
