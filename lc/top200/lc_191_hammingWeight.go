package main

import "fmt"

/*
给定一个数字，判断其二进制表达式中有多少个1
1、右移
2、与 1 &
3、累加计数
*/
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res += int(num & 1)
		num = num >> 1
	}
	return res
}

func hammingWeight1(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		res += int((num >> i) & 1)
	}
	return res
}

func main() {
	var num uint32 = 11111111111111111111111111111101
	fmt.Println(hammingWeight(num))
}
