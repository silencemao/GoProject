package main

import "fmt"

func countHomogenous(s string) int {
	s += " "
	const mod int = 1e9 + 7
	res, l, r, pre := 0, 0, 0, s[0]
	for ; r < len(s); r++ {
		if s[r] != pre {
			size := r - l
			res = res + (1+size)*size/2
			res %= mod
			l, pre = r, s[r]
		}
	}
	return res
}
func main() {
	s := "abbcccaa"
	fmt.Println(countHomogenous(s))
}
