package main

import "fmt"

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	pre, cur := 1, 2
	for i := 3; i <= n; i++ {
		pre, cur = cur, (pre+cur)%1000000007
	}
	return cur
}

func main() {
	n := 92
	fmt.Println(numWays(n))
}
