package main

import "fmt"

/*
买股票，股票可以买入卖出多次，计算可以获得的最大利润
可以贪心的去操作，只要今天价格比明天低，就可以今天买入，明天卖出，若后天价格更高，则可以明天买入，后天再卖出
*/
func maxProfit(prices []int) int {
	res := 0

	for i := 0; i <= len(prices) - 2; i++ {
		if prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
