package main

import (
	"fmt"
	"strconv"
)

/*
给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。

注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
*/
func multiply(num1 string, num2 string) string {
	if num1[0] == '0' || num2[0] == '0' {
		return ""
	}
	res := ""

	m, n := len(num1), len(num2)
	arr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			arr[i+j+1] += x * y
		}
	}
	fmt.Println(arr)
	for i := m + n - 1; i > 0; i-- {
		arr[i-1] += arr[i] / 10
		arr[i] %= 10
	}
	fmt.Println(arr)
	for i := 0; i < m+n; i++ {
		if i == 0 && arr[i] == 0 {
			continue
		}
		res += strconv.Itoa(arr[i])
	}
	return res
}
func main() {
	num1, num2 := "123", "456"
	fmt.Println(multiply(num1, num2))
}
