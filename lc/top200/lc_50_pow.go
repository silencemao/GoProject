package main

import (
	"fmt"
)

func powHelp(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else {
		if n%2 == 1 {
			return x * powHelp(x, n/2) * powHelp(x, n/2)
		}
		return powHelp(x, n/2) * powHelp(x, n/2)
	}
}

func pow(x float64, n int) float64 {
	flag := true
	if n < 0 {
		n = -n
		flag = false
	}
	res := powHelp(x, n)
	if !flag {
		res = 1.0 / res
	}
	return res
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	y := quickPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func pow1(x float64, n int) float64 {
	if n >= 0 {
		return quickPow(x, n)
	}
	return 1.0 / quickPow(x, -n)
}

func pow2(x float64, n int) float64 {
	res := 1.0
	flag := false
	if n < 0 {
		n = -n
		flag = true
	}
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		n >>= 1
		x *= x
	}
	if flag {
		res = 1.0 / res
	}
	return res
}

func main() {
	fmt.Println(pow2(2, 7))
}
