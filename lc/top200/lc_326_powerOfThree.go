package main

import "fmt"

/*
判断一个数是否为3的幂
除3再乘3是否为原数字 是则除3 继续判断， 否则返回false，
*/
func powerOfThree(n int) bool {
	for n > 0 {
		if n == 1 {
			return true
		}
		if n/3*3 != n {
			return false
		}
		n /= 3
	}
	return false
}

func main() {
	n := 0
	fmt.Println(powerOfThree(n))
}
