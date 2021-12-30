package main

import (
	"fmt"
	"sort"
)

/*
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:
可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
*/

func eraseOverlapIntervals(intervals [][]int) int {
	res := 0
	sort.Slice(intervals, func(i, j int) bool { // 按照右边界排序
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	pre := intervals[0][1]
	for i := 1; i < len(intervals); i++ { // 从左向右遍历
		if intervals[i][0] < pre { // 出现重叠
			res += 1
		} else { // 更新当前最大右边界
			pre = intervals[i][1]
		}
	}

	return res
}

func eraseOverlapIntervals1(intervals [][]int) int {
	res := 0
	sort.Slice(intervals, func(i, j int) bool { // 按照右边界排序
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	pre := intervals[len(intervals)-1][0]
	for i := len(intervals) - 2; i >= 0; i-- {
		if intervals[i][1] > pre {
			res += 1
		} else {
			pre = intervals[i][0]
		}
	}
	return res
}

func main() {
	intervals := [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}
	fmt.Println(eraseOverlapIntervals(intervals))
	fmt.Println(eraseOverlapIntervals1(intervals))
}
