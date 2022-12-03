package main

import "fmt"

/*
字符串中第2大的数字
*/
func secondHighest(s string) int {
	first, second := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			t := int(s[i] - '0')
			if t > first {
				first, second = t, first
			} else if t > second && t < first {
				second = t
			}
		}
	}
	if first == second {
		second = -1
	}
	return second
}
func main() {
	s := "ck077"
	fmt.Println(secondHighest(s))
}
