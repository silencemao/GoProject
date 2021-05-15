package main

import "fmt"

// 超时
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	first, second := 1, 2
	res := 0
	for i := 3; i <= n; i++ {
		res = first + second
		first, second = second, res
	}
	return res
}

func main() {
	n := 44
	fmt.Println(climbStairs(n))
	fmt.Println(climbStairs1(n))
}
