package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个字符串boxex = "110" boxes[i] = 1表示盒子i里面有一个小球，boxes[0] = 0表示盒子里面没有小球
返回一个列表 res res[i]表示将其它小球移动到第i个盒子里面需要的操作次数，一次一个小球只能移动到相邻盒子里面，一次只能操作一个小球
*/
func minOperation(boxes string) []int {
	res := make([]int, len(boxes))
	n := len(boxes)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != i && boxes[j] == '1' {
				res[i] += util.Abs(i - j)
			}
		}
	}
	return res
}

func minOperation2(boxes string) []int {
	res := make([]int, len(boxes))
	l, r := int(boxes[0]-'0'), 0
	operation := 0
	for i := 1; i < len(boxes); i++ {
		if boxes[i] == '1' {
			operation, r = operation+i, r+1
		}
	}
	res[0] = operation
	for i := 1; i < len(boxes); i++ {
		operation = operation + l - r
		if boxes[i] == '1' {
			l, r = l+1, r-1
		}
		res[i] = operation
	}

	return res
}

func main() {
	boxes := "001011"
	fmt.Println(minOperation(boxes))
	fmt.Println(minOperation2(boxes))
}
