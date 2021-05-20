package main

import "fmt"

// 可以买卖多次，只要买入之前先卖出就行。贪心的求解，只要今天的价钱比昨天高，我就可以卖出，若继续高，则累加相当于我后天卖出
func maxProfit3(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit3(nums))
}
