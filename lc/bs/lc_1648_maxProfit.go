package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
你有一些球的库存inventory，里面包含着不同颜色的球。一个顾客想要任意颜色 总数为orders的球。

这位顾客有一种特殊的方式衡量球的价值：每个球的价值是目前剩下的同色球的数目。比方说还剩下6个黄球，那么顾客买第一个黄球的时候该黄球的价值为6。
这笔交易以后，只剩下5个黄球了，所以下一个黄球的价值为5（也就是球的价值随着顾客购买同色球是递减的）
给你整数数组inventory，其中inventory[i]表示第i种颜色球一开始的数目。同时给你整数orders，表示顾客总共想买的球数目。你可以按照 任意顺序卖球。
请你返回卖了 orders个球以后 最大总价值之和。由于答案可能会很大，请你返回答案对 109+ 7取余数的结果。

arr[i]表示颜色为i的球的价格，也表示该颜色球的数量，当球的数量减少一个之后，价格也变为arr[i]-1
现在有人要买orders个球，如何售卖才能获得最大收益
*/
func maxProfit(inventory []int, orders int) int {
	res := 0
	mod := 1000000007

	for orders > 0 {
		sort.Ints(inventory)
		res = (res + inventory[len(inventory)-1]) % mod
		inventory[len(inventory)-1] -= 1
		orders -= 1
	}
	return res
}

func maxProfit1(inventory []int, orders int) int {
	maxItem, thresholdValue, count, res, mod := 0, -1, 0, 0, 1000000007
	for i := 0; i < len(inventory); i++ {
		if inventory[i] > maxItem {
			maxItem = inventory[i]
		}
	}
	low, high := 0, maxItem
	for low <= high {
		mid := low + ((high - low) >> 1)
		for i := 0; i < len(inventory); i++ {
			count += util.MaxInt(inventory[i]-mid, 0)
		}
		if count <= orders { // mid太大，取到的球个数不到orders个
			thresholdValue = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
		count = 0
	}
	count = 0
	for i := 0; i < len(inventory); i++ {
		count += util.MaxInt(inventory[i]-thresholdValue, 0)
	}
	count = orders - count
	for i := 0; i < len(inventory); i++ { //等差数列
		if inventory[i] >= thresholdValue {
			if count > 0 { // 取[th, inventory[i]]之间的数字 求和
				res += (thresholdValue + inventory[i]) * (inventory[i] - thresholdValue + 1) / 2
				count--
			} else { // 取[th+1, inventory[i]]之间的数字 求和
				res += (thresholdValue + 1 + inventory[i]) * (inventory[i] - thresholdValue) / 2
			}
		}
	}
	return res % mod
}

func maxProfit2(inventory []int, orders int) int {
	thresh, count, res, mod := 0, 0, 0, 1000000007
	sort.Ints(inventory)
	l, r := 0, inventory[len(inventory)-1]
	for l <= r {
		mid := l + (r-l)>>1
		for _, num := range inventory {
			count += util.MaxInt(num-mid, 0)
		}
		if count <= orders {
			thresh = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
		count = 0
	}
	count = 0
	for _, num := range inventory {
		count += util.MaxInt(num-thresh, 0)
	}
	count = orders - count

	for _, num := range inventory {
		if num >= thresh {
			if count > 0 {
				res += (num + thresh) * (num - thresh + 1) / 2
				count -= 1
			} else {
				res += (num + thresh + 1) * (num - thresh) / 2
			}
		}
	}

	return res % mod
}

func main() {
	inventory := []int{2, 8, 4, 10, 6}
	orders := 20
	//fmt.Println(maxProfit(inventory, orders))
	fmt.Println(maxProfit1(inventory, orders))
	fmt.Println(maxProfit2(inventory, orders))
}
