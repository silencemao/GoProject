package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
给你一个下标从 0开始的二维整数数组 events ，其中 events[i] = [startTimei, endTimei, valuei]。
第i个活动开始于startTimei，结束于endTimei，如果你参加这个活动，那么你可以得到价值 valuei 。你 最多可以参加两个时间不重叠活动，
使得它们的价值之和 最大。

请你返回价值之和的 最大值。

注意，活动的开始时间和结束时间是 包括在活动时间内的，也就是说，
你不能参加两个活动且它们之一的开始时间等于另一个活动的结束时间。更具体的，
如果你参加一个活动，且结束时间为 t ，那么下一个活动必须在 t + 1 或之后的时间开始。

*/
func maxTwoEvents(events [][]int) int {
	res := events[0][2]

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	for i := 1; i < len(events); i++ {
		if events[i][2] > res {
			res = events[i][2]
		}

		for j := 0; j < i; j++ {
			tmp := events[i][2]
			if events[j][1] < events[i][0] {
				tmp += events[j][2]
			}
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

func maxTwoEvents1(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	res := 0
	nextBig := make([]int, len(events))
	nextBig[len(events)-1] = events[len(events)-1][2]
	for i := len(events) - 2; i >= 0; i-- {
		nextBig[i] = util.MaxInt(nextBig[i+1], events[i][2])
	}

	for i := 0; i < len(events); i++ {
		sum := events[i][2]
		nexInd := bs(events, i+1, events[i][1])
		if nexInd < len(nextBig) {
			sum += nextBig[nexInd]
		}
		if sum > res {
			res = sum
		}
	}

	return res
}

func bs(events [][]int, l, end int) int {
	r := len(events)

	for l < r {
		mid := l + (r-l)>>1
		if events[mid][0] <= end {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	events := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}
	fmt.Println(maxTwoEvents(events))
	fmt.Println(maxTwoEvents1(events))
}
