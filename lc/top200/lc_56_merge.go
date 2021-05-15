package main

import (
	"fmt"
	"GoProject/leetcode/util"
	"sort"
)

func mergeInterval(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = util.MaxInt(intervals[i][1], res[len(res)-1][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func main() {
	intervals := [][]int{{1, 5}, {4, 7}, {3, 4}, {7, 9}, {10, 11}}
	res := mergeInterval(intervals)
	fmt.Println(res)
}
