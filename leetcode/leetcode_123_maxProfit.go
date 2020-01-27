package main

import "fmt"

/*
最大利润，至多进行两次交易，返回可以获得的最大利润
因为是最多两次交易，我们可以执行两次，即分别计算每次交易能获得的最大利润
第一次是先从左到右执行，即遍历一次数组，若当前price小于min，则更新min，这个min用于对后面的数值计算最大收益，
                                  若不小于min，则计算当前price卖出，min买入的收益，
同时用一个left数组存储以i结尾进行一次交易的收益

第二次我们倒着遍历数组，max为当前股票的最大价格，若当前价格大于max，则更新max，max用于对前面的数值计算最大收益
                                         若当前价格price小于max，则说明可以以当前价格price买入，max卖出，获得收益
同时用一个right数组存储以i为开始买入进行一次交易的收益
两次交易有没有冲突？不会的
以   {1, 6, 2, 2, 5, 10, 0}为例，
left=[0, 5, 5, 5, 5, 9, 9]
righ=[9, 8, 8, 8, 5, 0, 0]
分开来看，第一次交易取得最大值是9，而此时right是0
        第二次交易取得最大值是9，而此时left是0
而第一次交易取得5，第二次交易取得8时才是两次交易总和最大的时候
*/
func maxProfit123(prices []int) int {
	res := 0
	left := make([]int, len(prices))
	right := make([]int, len(prices))
	min := 1<<32
	localMax := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if localMax < prices[i] - min {
				localMax = prices[i] - min
			}
		}
		left[i] = localMax
	}
	max := -1000
	localMax = 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		} else {
			if localMax < max - prices[i] {
				localMax = max - prices[i]
			}
		}
		right[i] = localMax
	}
	for i := 0; i < len(prices); i++ {
		if res < left[i] + right[i] {
			res = left[i] + right[i]
		}
	}
	return res
}

func main() {
	prices := []int{1, 6, 2, 2, 10, 0}
	fmt.Println(maxProfit123(prices))
}
