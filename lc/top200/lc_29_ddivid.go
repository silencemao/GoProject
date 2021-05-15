package main

import "fmt"

/*
给定两个整数，不使用除法，计算两数相除的商
*/
// 相减，能减多少次，就说明商是多少
func divide(dividend int, divisor int) int {
	flag := true
	if dividend > 0 && divisor < 0 {
		flag = false
		divisor = -divisor
	} else if dividend < 0 && divisor > 0 {
		flag = false
		dividend = -dividend
	} else if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend == 0 {
		return 0
	}
	res := 0
	for dividend >= divisor {
		if (flag && res == 1<<31-1) || (!flag && -res == 1<<31) {
			return 1<<31 - 1
		}
		res += 1
		dividend -= divisor
	}
	if !flag {
		res = -res
	}
	return res
}

/*
既然相减可以计算得到结果，但是每次都减divisor是不是有点慢，那能不能一次减n倍的divisor是不是也可以，
反过来当divisor + divisor < dividend的时候，dividend至少是divisor的2倍，cur += cur
那我们对divisor翻倍 若 仍满足divisor+divisor < dividend dividend至少是divisor的4倍， cur += 2
*/
func divide1(dividend, divisor int) int {
	flag := true
	if dividend < 0 {
		flag = !flag
		dividend = -dividend
	}
	if divisor < 0 {
		flag = !flag
		divisor = -divisor
	}
	if dividend == 0 || divisor > dividend {
		return 0
	}

	res := 0
	//8ms
	//for divisor <= dividend {
	//	cur := divisor
	//	multip := 1
	//	for cur+cur <= dividend {
	//		cur += cur
	//		multip += multip
	//	}
	//	dividend -= cur
	//	res += multip
	//}
	// 0ms
	for divisor <= dividend {
		cur, multip := divisor, 1
		for cur<<1 <= dividend {
			cur, multip = cur<<1, multip<<1
		}
		res += multip
		dividend -= cur
	}

	if -res < -1<<31 || res >= 1<<31-1 {
		return 1<<31 - 1
	}
	if !flag {
		res = -res
	}

	return res
}

func main() {
	fmt.Println(divide1(-1, 1))
}
