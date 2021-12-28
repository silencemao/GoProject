package main

import "fmt"

/*
在柠檬水摊上，每一杯柠檬水的售价为5美元。顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。

每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
注意，一开始你手头没有任何零钱。
给你一个整数数组 bills ，其中 bills[i] 是第 i 位顾客付的账。如果你能给每位顾客正确找零，返回true，否则返回 false。

对于bill[i] = 5直接收下，bills[i] = 10找一个5 bills[i] = 20优先给1个5 1个10 若10不够，则再给两个5
*/
func lemonadeChange(bills []int) bool {
	size := len(bills)
	if size < 1 {
		return true
	}
	if bills[0] != 5 {
		return false
	}
	cnt5, cnt10 := 0, 0
	for _, v := range bills {
		if v == 5 {
			cnt5 += 1
		} else if v == 10 {
			cnt5 -= 1
			cnt10 += 1
		} else {
			cnt5 -= 1
			if cnt10 > 0 {
				cnt10 -= 1
			} else {
				cnt5 -= 2
			}

		}
		if cnt5 < 0 || cnt10 < 0 {
			return false
		}
	}
	return true
}
func main() {
	bills := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	fmt.Println(lemonadeChange(bills))
}
