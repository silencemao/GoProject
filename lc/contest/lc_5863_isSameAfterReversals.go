package main

import "fmt"

func isSameAfterReversals(num int) bool {
	if num < 10 {
		return true
	}
	last := num % 10
	if last != 0 {
		return true
	}
	return false
}

func main() {
	num := 675
	fmt.Println(isSameAfterReversals(num))
}
