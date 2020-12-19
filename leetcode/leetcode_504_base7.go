package main

import (
	"fmt"
	"strconv"
)

/*
给定10进制的数，转换为7进制表示的字符串
*/

func convertToBase7(num int) string {
	res := ""
	for num != 0 {
		res = strconv.Itoa(num%7) + res
		num = num / 7
	}
	if num < 0 {
		return "-" + res
	}
	return res
}

func main() {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(4))
	fmt.Println(convertToBase7(-7))

}
