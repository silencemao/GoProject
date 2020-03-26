package main

import "fmt"

/*
不使用乘法和除法、取余 完成 两数相除的运算

不能使用除法和减法，就只能使用 位运算，
*/
func divide(dividend, divisor int) int {
	maxValue := 1<<31 - 1
	minValue := -1<<31

	if dividend==0 {
		return 0
	}
	if dividend==minValue && divisor==-1 {
		return maxValue
	}

	flag := (dividend>0)==(divisor>0)

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := 0
	for dividend >= divisor {
		t := divisor
		p := 1
		for dividend >= (t<<1) {
			t <<= 1
			p <<= 1
		}
		res += p
		dividend -= t
	}
	if !flag {
		res = -res
	}
	return res
}

func divide1(dividend, divisor int) int {

	maxValue := 1<<31 - 1
	minValue := -1<<31

	if dividend==0 {
		return 0
	}
	if dividend==minValue && divisor==-1 {
		return maxValue
	}

	flag := (dividend>0)==(divisor>0)

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	count := 0
	for dividend >= divisor {
		dividend -= divisor
		count++
	}
	if !flag {
		count = -count
	}
	return count
}

func main() {
	dividend := -2147483648
	divisor := -1
	fmt.Println(divide(dividend, divisor), divide1(dividend, divisor))
}
