package main

import (
	"fmt"
	"strings"
)

/*
atoi 字符串转为数字，注意范围 小于等于-1<<31时直接返回-1<<31 大于等于1<<31-1时直接返回1<<31-1
*/
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) <= 0 {
		return 0
	}

	res := 0
	bs := []byte(s)
	if bs[0] != '-' && bs[0] != '+' && (bs[0] < '0' || bs[0] > '9') {
		return 0
	}
	isNative := false
	if bs[0] == '-' {
		isNative = true
	}
	for i := 0; i < len(bs); i++ {
		if i == 0 && (isNative || bs[0] == '+') {
			continue
		}
		if bs[i] < '0' || bs[i] > '9' {
			break
		}
		tmp := int(bs[i] - '0')
		if !isNative {
			if res > ((1<<31-1)-tmp)/10 {
				return 1<<31 - 1
			}
		} else {
			if -res < (-(1<<31)+tmp)/10 {
				return -(1 << 31)
			}
		}
		res = res*10 + int(bs[i]-'0')
	}
	if isNative {
		res = -res
	}
	return res
}

func main() {
	s := "-2147483649"
	fmt.Println(myAtoi(s))
}
