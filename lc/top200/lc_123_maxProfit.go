package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 最多买入两次
func maxProfit2Time(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	left, right := make([]int, len(prices)), make([]int, len(prices))

	curMin := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < curMin {
			curMin = prices[i]
		}
		left[i] = util.MaxInt(left[i-1], prices[i]-curMin)
	}

	curMax := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > curMax {
			curMax = prices[i]
		}
		right[i] = util.MaxInt(right[i+1], curMax-prices[i])
	}
	res := 0
	for i := 0; i < len(prices); i++ {
		if left[i]+right[i] > res {
			res = left[i] + right[i]
		}
	}
	return res
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit2Time(prices))
}
