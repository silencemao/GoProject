package main

import "fmt"

const pick = 10

func guess(num int) int {
	if num == pick {
		return 0
	} else if num < pick {
		return -1
	} else {
		return 1
	}
}

func guessNum(n int) int {
	if guess(n) == 0 {
		return n
	}
	left, right := 1, n
	for left < right {
		mid := (left + right) >> 1
		t := guess(mid)
		if t == 0 {
			return mid
		} else if t < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	fmt.Println(guessNum(6))
}
