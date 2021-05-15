package main

import "fmt"

const(
	M = 3
)

func combine(a []int, b []int, n, m int) {
	for i := n; i >= m; i-- {
		b[m-1] = i - 1
		if m > 1 {
			combine(a, b, i-1, m-1)
		} else {
			for j := M-1; j >= 0; j-- {
				fmt.Print(a[b[j]], " ")
			}
			fmt.Println()
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3)
	combine(a, b, len(a), 3)
}
