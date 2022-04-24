package main

import "fmt"

/*
给你一个非负整数 x ，计算并返回x的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/
func mySqrt(x int) int {
	l, r := 0, x/2+1
	mid := 0
	for l <= r {
		mid = l + (r-l)>>1
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
func main() {
	x := 8

	fmt.Println(mySqrt(x))
}
