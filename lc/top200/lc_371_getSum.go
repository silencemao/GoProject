package main

import "fmt"

func getSum(a int, b int) int {
	res := a ^ b          // 无进位加法
	carry := (a & b) << 1 // 进位
	for carry != 0 {      // 此处要注意进位
		res, carry = res^carry, (res&carry)<<1
	}
	return res
}

func main() {
	a, b := -1, -21
	fmt.Println(getSum(a, b))
}
