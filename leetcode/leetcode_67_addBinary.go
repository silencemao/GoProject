package main

import (
	"fmt"
	"strconv"
)

/*
给定两个字符串a b，均是二进制表示的值，返回a b相加之后的二进制表示
*/
func addBinary(a, b string) string {
	as := []byte(a)
	bs := []byte(b)
	res := ""
	m, n, carry := len(as) -1, len(bs) - 1, 0
	for m >= 0 || n >= 0 {
	    p, q := 0, 0
		if m >= 0 {
		    p = int(as[m] - '0')
			m--
		}
		if n >= 0 {
		    q = int(bs[n] - '0')
			n--
		}
		sum := p + q + carry
		res = strconv.Itoa(sum % 2) + res
		carry = sum / 2
	}
	if carry == 1 {
	    res = "1" + res 
	}
	return res 
}

func main() {
    a, b := "11", "1"
	fmt.Println(addBinary(a, b))
}
