package main

import "fmt"

/*
最大利润，只能买入一次卖出一次
*/
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxProfit1(prices []int) int {
	res := 0
	for i := 0; i < len(prices) - 1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				res = Max(res, prices[j] - prices[i])
			}
		}
	}
	return res
}
/*
一次遍历 累加求和
*/
func maxProfit2(prices []int) int {
	res := 0
	curMax := 0
	for i := 1; i < len(prices); i++ {
		curMax = Max(0, curMax + prices[i] - prices[i-1])
		res = Max(curMax, res)
	}
	return res
}
/*
一次遍历，记录最小值，计算当前值与最小值之间的差值，然后更新较大的差值
*/
func maxProfit3(prices []int) int {
	res, buy := 0, 1<<32
	for _, price := range prices {
		buy = Min(buy, price)
		res = Max(res, price - buy)
	}
	return res
}
func main() {
	prices := []int{0, 6, -3, 7}
	fmt.Println(maxProfit1(prices), maxProfit2(prices), maxProfit3(prices))
}
