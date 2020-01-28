package main

import (
	"fmt"
	"math"
)
/*
买入卖出的最大收益，中间需有一个cooldown，即卖出之后必须停止交易一天，才能进行下次交易
profit代表截止到i天的收益，profit[i]表示第i天为cooldown的总收益，或者卖出的总收益，为什么不能表示买入的收益？因为买入之后必须要卖出才能有收益，当i为
最后一天时，没有机会卖出了，肯定是白买的，所以profit[i]只代表截止到第i天为cooldown或者卖出的总收益
若profit[i]为cooldown，则profit[i]=profit[i-1]
若profit[i]为卖出，则需要确定买入的时候，以及买入之前的总收入。
买入的时候为0-i的某一天，用j表示，即在price[j]买入，profit[i]卖出。买入之前的总收入如何获得的呢？
price[j]买入，则profit[j-1]为cooldown，profit[j-1]的收益为profit[j-2]
所以profit[i]=profit[j-2] + price[i] - price[j]

因此profit[i] = max(profit[i-1], profit[j-2] + price[i] - price[j])
*/
func maxProfit309(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	n := len(prices)
	profit := make([]int, n)
	for i := 0; i < n; i++ {
		if i==0 {
			profit[i] = 0
		} else if i==1 {
			profit[i] = int(math.Max(0, float64(prices[i]-prices[i-1])))
		} else {
			profit[i] = profit[i-1]
			for j := 0; j < i; j++ {
				prevProfit := 0
				if j > 2 {
					prevProfit = profit[j - 2]
				}
				profit[i] = int(math.Max(float64(profit[i]), float64(prevProfit+prices[i]-prices[j])))
			}
		}
	}

	return profit[n-1]
}
/*
s0代表休息后休息 或 卖出后休息
s1代表休息后买入 或 买入后休息
s2代表休息后卖出 或 买入后卖出
*/
func maxProfit309_1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	s0, s1, s2 := 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		pre0, pre1, pre2 := s0, s1, s2
		s0 = int(math.Max(float64(pre0), float64(pre2)))
		s1 = int(math.Max(float64(pre0-prices[i]), float64(pre1)))
		s2 = pre1 + prices[i]
	}
	return int(math.Max(float64(s0), float64(s2)))
}

func main() {
	prices := []int{1, 2, 6, 0, 10}
	fmt.Println(maxProfit309(prices), maxProfit309_1(prices))
}
