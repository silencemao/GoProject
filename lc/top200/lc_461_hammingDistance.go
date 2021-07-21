package main

import "fmt"

/*
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
todo: 若存在负数如何处理
*/
func hammingDistance(x int, y int) int {
	res := 0
	for x > 0 || y > 0 {
		res += (x & 1) ^ (y & 1)
		x = x >> 1
		y = y >> 1
	}

	return res
}
func main() {
	x, y := 1, 3
	fmt.Println(hammingDistance(x, y))
}
