package main

import (
	"fmt"
	"math"
)

/*
给定一个数c，判断是否存在两个整数 a^2+b^2=c，a b可以相等
*/
/*双指针*/
func judgeSquareSum(c int) bool {
	l, r := 0, c/2+1
	for l <= r { // 可以存在两个相等的整数
		tmp := l*l + r*r
		if tmp == c {
			return true
		}
		if tmp > c {
			r -= 1
		} else {
			l += 1
		}
	}
	return false
}

/*
二分法
a的范围是0-sqrt(c)
b的范围是a-sqrt(c)
*/
func judgeSquareSum1(c int) bool {
	x := int(math.Sqrt(float64(c)))
	for i := 0; i < x+1; i++ {
		target := c - i*i
		l, r := i, x
		for l <= r {
			mid := l + (r-l)>>1
			tmp := mid * mid
			if tmp == target {
				return true
			}
			if tmp > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}

func main() {
	c := 130
	fmt.Println(judgeSquareSum1(c))
	fmt.Println(judgeSquareSum(c))
}
