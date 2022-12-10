package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

func maxHeight(cuboids [][]int) int {
	res, n := 0, len(cuboids)
	for _, cub := range cuboids {
		sort.Ints(cub)
	}

	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]
		return a[0] < b[0] || a[0] == b[0] && (a[1] < b[1] || a[1] == b[1] && a[2] < b[2])
	})

	dp := make([]int, len(cuboids))
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2]
		for j := i - 1; j >= 0; j-- {
			if cuboids[i][0] >= cuboids[j][0] && cuboids[i][1] >= cuboids[j][1] && cuboids[i][2] >= cuboids[j][2] {
				dp[i] = util.MaxInt(dp[i], dp[j]+cuboids[i][2])
			}
		}
		res = util.MaxInt(res, dp[i])
	}
	return res
}

func main() {
	cuboids := [][]int{{95, 37, 53}, {45, 23, 12}, {50, 45, 20}}
	fmt.Println(maxHeight(cuboids))
}
