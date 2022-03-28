package main

import (
	"fmt"
	"strconv"
)

/*
给你一个仅由数字组成的字符串 s 。
请你判断能否将 s 拆分成两个或者多个 非空子字符串 ，使子字符串的 数值 按 降序 排列，且每两个 相邻子字符串 的数值之 差 等于 1 。
例如，字符串 s = "0090089" 可以拆分成 ["0090", "089"] ，数值为 [90,89] 。这些数值满足按降序排列，且相邻值相差 1 ，这种拆分方法可行。
另一个例子中，字符串 s = "001" 可以拆分成 ["0", "01"]、["00", "1"] 或 ["0", "0", "1"] 。然而，所有这些拆分方法都不可行，因为对应数值分别是 [0,1]、[0,1] 和 [0,0,1] ，都不满足按降序排列的要求。
如果可以按要求拆分 s ，返回 true ；否则，返回 false 。

将字符串s拆分,转换为数字nums = [....]，若nums[i]-nums[i+1] = 1，返回true，不满足返回false。

1、cnt记录拆分的数字个数
2、返回条件 当指针指到s末尾时，cnt>1
3、当cnt=0或prev-num=1，则继续递归
*/
func help18492(s string, prev int, cnt *int, ind int) bool {
	if ind >= len(s) {
		return *cnt > 1
	}
	for i := ind; i < len(s); i++ {
		num, _ := strconv.Atoi(s[ind : i+1])
		if *cnt == 0 || prev-num == 1 {
			*cnt += 1
			if help18492(s, num, cnt, i+1) {
				return true
			}
			*cnt -= 1
		}
	}
	return false
}

func splitString(s string) bool {
	cnt := 0
	m := help18492(s, 0, &cnt, 0)
	return m
}

func main() {
	s := "4771447713"
	fmt.Println(splitString(s))
}
