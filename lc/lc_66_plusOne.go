package main

import "fmt"

/*
给定一个数组，如[4, 3, 2, 1]数组可以表示整数为4321，数组最后一个数字加1，得到4322，返回加1之后的数组[4, 3, 2, 2]
*/

func plusOne(digits []int) []int {
	if len(digits)==0 {
		return digits
	}

	var num int64
	for _, n := range digits {
		num = num * 10 + int64(n)
	}

	num += 1
	digits = nil
	for num != 0 {
		digits = append(digits, int(num % 10))
		num = num / 10
	}
	digits = Reverse(digits)
	return digits
}

func Reverse(s []int) []int {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func plusOne1(digits []int) []int {

	lastPlusOne := digits[len(digits) - 1] + 1
	if lastPlusOne < 10 {
		digits[len(digits) - 1] = lastPlusOne
		return digits
	}
	digits[len(digits) - 1] = lastPlusOne % 10
	tmp := lastPlusOne / 10

	for i := len(digits) - 2; i >= 0; i-- {
		num := digits[i] + tmp
		if num < 10 {
			digits[i] = num
			tmp = 0
			break
		} else {
			remain := num % 10
			digits[i] = remain
			tmp = num / 10
		}
	}
	if tmp != 0 {
		digits = append([]int{tmp}, digits...)
	}

	return digits
}

func main() {
	a := []int{4, 3, 2, 1}
	fmt.Println(a)
	fmt.Println(plusOne(a))

	fmt.Println(a)
	fmt.Println(plusOne1(a))
}
