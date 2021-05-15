package main

import "fmt"

// 给定一个数组，返回加一之后的数组
// 注意 999+1=1000这种情况，会多出来一个数
func plusOne(digits []int) []int {
	if len(digits) < 1 {
		return digits
	}
	if digits[0] == 0 {
		return []int{1}
	}
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if i == len(digits)-1 {
			sum += 1
		}
		digits[i] = sum % 10
		carry = sum / 10
		if carry == 0 {
			break
		}
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{1, 9, 9, 9, 9}
	fmt.Println(plusOne(digits))
}
