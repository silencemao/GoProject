package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	i := 0
	for ; i <= x/2; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return i - 1
}

// 二分法
func binaryHelp(l, r, x int) int {
	for l <= r {
		mid := (l + r) / 2
		tmp := mid * mid
		if tmp == x {
			return mid
		}
		if tmp > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l*l > x {
		return l - 1
	}
	return l
}

func mySqrt1(x int) int {
	l, r := 0, x/2
	res := binaryHelp(l, r, x)
	return res
}

func main() {
	x := 0
	fmt.Println(mySqrt1(x))
}
