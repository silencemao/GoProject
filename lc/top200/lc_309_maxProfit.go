package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

// 带有冷冻期的 可以进行多次买卖 卖出之后必须冷冻一天才能卖出
func maxProfit4(prices []int) int {
	res := 0
	if len(prices) == 1 {
		return res
	}
	s0, s1, s2 := 0, -prices[0], 0
	for i, p := range prices {
		if i == 0 {
			continue
		}
		pre0, pre1, pre2 := s0, s1, s2
		s0 = util.MaxInt(pre2, pre0)
		s1 = util.MaxInt(pre0-p, pre1)
		s2 = p + pre1
	}
	return util.MaxInt(s0, s2)
}

func main() {
	nums := []int{6, 1, 6, 4, 3, 0, 2}
	fmt.Println(maxProfit4(nums))
}
