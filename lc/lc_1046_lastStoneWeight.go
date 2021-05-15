package main

import (
	"fmt"
	"sort"
)

/*
给定一个数组，从中选出最大的两个数x，y(x<=y)，若x==y，则x y都从数组中删除，
若x!=y，则删除x，y = y - x
*/
func lastStoneWeight(stones []int) int {
	res := 0

	for len(stones) > 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		y, x := stones[0], stones[1]
		if x==y {
			stones = stones[2:]
		} else {
			stones[1] = y - x
			stones = stones[1:]
		}
	}

	if len(stones)==1 {
		res = stones[0]
	}

	return res
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}