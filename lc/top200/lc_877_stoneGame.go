package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func help877_1(piles []int, cur0, cur1, left, right, player int) bool {
	if left > right {
		return cur0 > cur1
	}
	if player == 0 {
		return help877_1(piles, cur0+piles[left], cur1, left+1, right, 1) || help877_1(piles, cur0+piles[right], cur1, left, right-1, 1)
	} else {
		return help877_1(piles, cur0, cur1+piles[left], left+1, right, 0) || help877_1(piles, cur0, cur1+piles[right], left, right-1, 0)
	}
}

// 超时
func stoneGame(piles []int) bool {
	return help877_1(piles, 0, 0, 0, len(piles)-1, 0)
}

func stoneGame2(piles []int) bool {
	size := len(piles)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = piles[i]
	}
	for i := size - 2; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			dp[i][j] = util.MaxInt(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][size-1] > 0
}

func main() {
	stones := []int{5, 3, 4, 5}
	fmt.Println(stoneGame2(stones))
}
