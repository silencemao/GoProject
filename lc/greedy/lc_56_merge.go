package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*/
func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = util.MaxInt(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
